# gRPC

API gRPC em Go para categorias e cursos — contrato em Protocol Buffers, camada de serviço e persistência.

---

## Este projeto

Este repositório contém uma implementação em **Go** de um serviço gRPC para **categorias e cursos**: contrato da API em Protocol Buffers, camada de persistência e implementação do método **CreateCategory**.

### Para que serve

- **Expor uma API gRPC** de categorias: no contrato (`.proto`) está definido o `CategoryService` com o método **CreateCategory** (chamada unária). Outros métodos e o serviço de cursos podem ser estendidos a partir da mesma base.
- **Persistir dados**: a camada em `internal/database` implementa o acesso a tabelas **categories** e **courses** (PostgreSQL-compatível). A camada de serviço em `internal/service` usa o banco e implementa o handler gRPC.

### Como funciona

1. **Contrato (proto)** — Em `proto/course_category.proto` estão as mensagens (`Category`, `CreateCategoryRequest`, `CategoryResponse`) e o serviço `CategoryService` com `CreateCategory`. O código Go é gerado com `protoc` em `internal/pb` (veja [Instalando compilador e plugins](#instalando-compilador-e-plugins)).
2. **Camada de dados** — Em `internal/database`: **Category** (`Create`, `FindAll`, `FindByCourseID`) e **Course** (`Create`, `FindAll`, `FindByCategoryID`).
3. **Camada de serviço** — Em `internal/service/category.go` o `CategoryService` implementa `CreateCategory`: recebe a requisição gRPC, chama o banco e devolve a categoria criada.

### Implementando CreateCategory

A implementação fica em `internal/service/category.go` e conecta o contrato gRPC (`internal/pb`) à persistência (`internal/database`).

- **`CategoryService`** — struct que implementa `CategoryServiceServer`, embute `pb.UnimplementedCategoryServiceServer` e recebe `database.Category`.
- **`NewCategoryService(categoryDB)`** — construtor; o serviço é registrado no `grpc.Server` com `pb.RegisterCategoryServiceServer(grpcServer, categoryService)`.
- **Fluxo de `CreateCategory`:** servidor recebe `*pb.CreateCategoryRequest` → chama `CategoryDB.Create(req.Name, req.Description)` → retorna `*pb.CategoryResponse` com a categoria (id, name, description) ou repassa o erro.

### Criando CategoryList no proto

No arquivo `proto/course_category.proto` foram adicionadas a mensagem **`CategoryList`** e o RPC **`ListCategories`** para permitir que o cliente obtenha a lista de todas as categorias em uma única chamada.

**O que foi adicionado no `.proto`:**

- **`message blank {}`** — Mensagem vazia usada como tipo de requisição quando o RPC não precisa de parâmetros de entrada (só “listar tudo”).
- **`message CategoryList { repeated Category categories = 1; }`** — Mensagem que agrupa várias categorias. No Protocol Buffers, a resposta de um RPC precisa ser uma única mensagem; para devolver uma **lista** de categorias, usa-se uma mensagem que contém um campo `repeated Category`.
- **`rpc ListCategories (blank) returns (CategoryList);`** — Método que recebe nada (ou um `blank`) e retorna um `CategoryList`, ou seja, a lista de categorias cadastradas.

**Motivação da implementação:**

- **Listar categorias via API:** O cliente precisa de uma operação “listar todas as categorias” (por exemplo para montar um menu, um filtro ou uma tela de administração). O `ListCategories` expõe exatamente isso no contrato gRPC.
- **Contrato explícito:** Definir `CategoryList` no `.proto` deixa o formato da resposta documentado e tipado: o cliente sempre recebe uma mensagem com o campo `categories` (array de `Category`). Isso evita ambiguidade e facilita a evolução da API (por exemplo, no futuro adicionar metadados como `total_count` na mesma mensagem).
- **Reuso da mensagem `Category`:** A lista reutiliza a mesma mensagem `Category` já usada em `CreateCategory`, mantendo o contrato consistente e evitando duplicação.
- **Alinhamento com a camada de dados:** A camada `internal/database` já oferece `FindAll()` para categorias. O `ListCategories` no serviço faz a ponte entre esse método e o contrato gRPC, devolvendo as categorias no formato `CategoryList`.

Depois de alterar o `.proto`, é necessário regenerar o código com `protoc --go_out=. --go-grpc_out=. proto/course_category.proto` e implementar o método `ListCategories` em `internal/service/category.go` (chamando `CategoryDB.FindAll()` e montando o `CategoryList`).

### Implementando ListCategories

A implementação do RPC **ListCategories** está em `internal/service/category.go`. Ela atende chamadas que não enviam parâmetros (tipo `*pb.Blank`) e devolve todas as categorias persistidas no banco no formato `*pb.CategoryList`.

**Fluxo do método:**

1. O servidor gRPC chama `ListCategories(ctx, req)` com `req` do tipo `*pb.Blank` (sem uso no método; o contrato exige uma mensagem de entrada).
2. O serviço chama `c.CategoryDB.FindAll()`, que consulta a tabela `categories` e retorna um slice de `database.Category` ou erro.
3. O serviço converte cada `database.Category` em `*pb.Category` (id, name, description) e monta um slice `categoriesResponse`.
4. Retorna `&pb.CategoryList{ Categories: categoriesResponse }`. Em caso de erro do banco, o erro é repassado e o gRPC devolve o status apropriado ao cliente.

**Motivação da implementação:**

- **Expor o FindAll via gRPC:** A camada de dados já oferece `FindAll()` para categorias. O `ListCategories` é a porta de entrada gRPC para essa operação: o cliente chama um único RPC e recebe a lista completa, sem precisar de filtros ou paginação (que podem ser adicionados depois com novas mensagens/RPCs).
- **Resposta tipada e consistente:** A resposta é sempre `CategoryList` com o campo `categories` (repeated `Category`), o que mantém o contrato estável e previsível para qualquer cliente (Evans, outro serviço, app mobile).
- **Reuso da mesma camada:** Tanto `CreateCategory` quanto `ListCategories` usam o mesmo `CategoryDB`; a única diferença é a operação (Create vs FindAll) e o mapeamento do resultado para as mensagens do proto (`*pb.Category` ou `*pb.CategoryList`).

Com isso, o `CategoryService` passa a oferecer criação e listagem de categorias; no Evans, após `package pb` e `service CategoryService`, é possível chamar tanto `CreateCategory` quanto `ListCategories`.

### Buscando uma categoria (GetCategory)

Foi adicionada a operação de **buscar uma categoria pelo ID**: no `.proto` a mensagem **`GetCategoryRequest`** e o RPC **`GetCategory`**, na camada de dados o método **`FindByID`** e no serviço o handler que liga um ao outro.

**O que foi adicionado no `.proto`:**

- **`message GetCategoryRequest { string id = 1; }`** — Requisição com o identificador da categoria a ser buscada.
- **`rpc GetCategory (GetCategoryRequest) returns (Category);`** — Método que recebe o ID e retorna uma única `Category` (ou erro se não existir).

**Implementação:**

- **Camada de dados** (`internal/database/category.go`): **`FindByID(id string) (Category, error)`** — Executa `SELECT id, name, description FROM categories WHERE id = $1` e devolve a categoria ou erro (ex.: `sql.ErrNoRows` se o ID não existir).
- **Camada de serviço** (`internal/service/category.go`): **`GetCategory(ctx, req)`** — Recebe `*pb.GetCategoryRequest`, chama `c.CategoryDB.FindByID(req.Id)`, converte o resultado em `*pb.Category` e retorna. Se o banco retornar erro (categoria não encontrada), o erro é repassado e o gRPC traduz em status apropriado (ex.: `NotFound`) para o cliente.

**Motivação da implementação:**

- **Consulta por ID:** O cliente precisa obter os dados de **uma** categoria conhecida (tela de detalhe, formulário de edição, validação antes de vincular a um curso). O `GetCategory` expõe essa consulta de forma explícita no contrato.
- **Contrato claro:** A requisição é tipada (`GetCategoryRequest` com `id`); a resposta é a mensagem `Category` já usada em criação e listagem. Isso mantém a API consistente e documentada.
- **Reuso da camada de dados:** O `FindByID` no banco é a única consulta por chave; o serviço apenas delega e mapeia para o tipo do proto, no mesmo padrão de `CreateCategory` e `ListCategories`.

No Evans, após `package pb` e `service CategoryService`, use `call GetCategory` e informe o campo `id` (ex.: um UUID retornado por `CreateCategory` ou `ListCategories`) para testar a busca.

### Criando servidor gRPC

O ponto de entrada do servidor está em **`cmd/grpcServer/main.go`**. Ele sobe o gRPC na porta **50051** e registra o `CategoryService` com reflexão habilitada (útil para ferramentas como `grpcurl` ou **Evans**).

**Ordem da implementação:**

1. **Conexão com o banco** — `sql.Open("sqlite3", "db.sqlite")` (ou troque para PostgreSQL conforme a camada em `internal/database`). O `defer db.Close()` garante o fechamento ao encerrar o processo.

2. **Camada de dados** — Cria-se o repositório de categorias: `categoryDB := database.NewCategory(db)`.

3. **Camada de serviço** — Cria-se o serviço que implementa o gRPC: `categoryService := service.NewCategoryService(*categoryDB)`. Esse é o tipo que implementa `CategoryServiceServer` e será registrado no servidor.

4. **Servidor gRPC** — `grpc.NewServer()` cria o servidor; `pb.RegisterCategoryServiceServer(grpcServer, categoryService)` registra o `CategoryService`; `reflection.Register(grpcServer)` ativa a reflexão para descoberta de serviços.

5. **Listen e Serve** — `net.Listen("tcp", ":50051")` abre a porta; `grpcServer.Serve(lis)` bloqueia e atende as chamadas RPC (por exemplo, `CreateCategory`).

**Como rodar:** na raiz do projeto: `go run cmd/grpcServer/main.go`. O servidor fica ouvindo em `localhost:50051` até ser interrompido.

### Rodando o projeto e testando com Evans

**[Evans](https://github.com/ktr0731/evans)** é um cliente gRPC em modo REPL (linha de comando interativa). Ele usa a **reflexão** do servidor para listar serviços e métodos e permite chamar RPCs sem escrever código cliente. Útil para testar a API logo após subir o servidor.

**1. Instalar o Evans**

```sh
go install github.com/ktr0731/evans@latest
```

**Motivo:** O Evans fica em `$(go env GOPATH)/bin`. Garanta que esse diretório esteja no `PATH` (ex.: `export PATH="$PATH:$(go env GOPATH)/bin"`).

**2. Subir o servidor gRPC**

Em um terminal, na raiz do projeto:

```sh
go run cmd/grpcServer/main.go
```

**Motivo:** O servidor precisa estar rodando na porta **50051** para o Evans (ou qualquer cliente) se conectar e chamar `CreateCategory`.

**3. Criar o banco e a tabela (primeira vez)**

O servidor usa SQLite com o arquivo `db.sqlite`. Se o arquivo ou a tabela não existir, crie:

```sh
sqlite3 db.sqlite
```

Dentro do SQLite:

```sql
create table categories (id string, name string, description string);
```

**Motivo:** A camada `internal/database` espera a tabela `categories` com as colunas `id`, `name` e `description`. Sem ela, `CreateCategory` falha ao persistir.

**4. Abrir o Evans em modo REPL**

Em **outro** terminal (com o servidor ainda rodando):

```sh
evans -r repl
```

**Motivo:** `-r` usa **reflexão** no servidor (por isso o `reflection.Register(grpcServer)` no `main.go`). O Evans descobre sozinho os pacotes, serviços e métodos disponíveis em `localhost:50051`.

**5. No Evans, usar nesta ordem**

| Comando | O que faz |
|---------|-----------|
| `package pb` | Seleciona o pacote do `.proto` (definido em `option go_package` / package no proto). Sem isso, o Evans não sabe em qual serviço atuar. |
| `service CategoryService` | Seleciona o serviço `CategoryService`. Sem isso, o comando `call` não sabe qual RPC usar. |
| `call CreateCategory` | Chama o RPC `CreateCategory`; o Evans pede os campos da requisição (name, description) e exibe a resposta. |

**Motivo:** O Evans exige que **package** e **service** estejam selecionados antes de `call`. Caso contrário aparecem erros como *"package unselected"* ou *"service unselected"*.

**Resumo:** Instale o Evans → suba o servidor → crie a tabela se necessário → em outro terminal rode `evans -r repl` → `package pb` → `service CategoryService` → `call CreateCategory`.

### Alterações no contrato (.proto): regenerar código

Sempre que você **alterar** o arquivo `proto/course_category.proto` (novas mensagens, novos métodos, novos campos), é necessário **regenerar** o código Go para que `internal/pb` e o servidor/serviço continuem alinhados com o contrato. Na raiz do projeto execute:

```sh
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

**Motivo:** Os arquivos em `internal/pb/*.pb.go` são gerados pelo `protoc`. Eles não são atualizados automaticamente; qualquer mudança no `.proto` exige rodar o comando de novo. Depois, ajuste a implementação em `internal/service` (e no `main` do servidor, se houver novos serviços) conforme o novo contrato.

### Motivação

- Praticar gRPC em Go com contrato primeiro (`.proto`) e código gerado.
- Separar contrato (proto), persistência (database) e serviço (service).
- Base para evoluir com mais métodos e tipos de chamada (streaming).

---

## Instalando compilador e plugins

Para gerar código gRPC a partir dos arquivos `.proto`, é necessário o compilador **Protocol Buffers** (`protoc`) e os **plugins** Go.

### Pré-requisitos

- **Go** — uma das duas últimas versões principais ([Guia de instalação do Go](https://go.dev/doc/install)).
- **protoc** — versão 3 ([Protoc Installation](https://protobuf.dev/programming-guides/protoc-installation/)).
- **Plugins Go:**
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
  Inclua `$(go env GOPATH)/bin` no `PATH` (ex.: `export PATH="$PATH:$(go env GOPATH)/bin"`).

### Gerando código a partir do `.proto`

No diretório raiz do projeto:

1. **Recarregar o ambiente** (se precisar que o shell enxergue `protoc` e os plugins):
   ```sh
   source ~/.bashrc
   ```
   **O que faz:** Reaplica o `~/.bashrc` no shell atual (ex.: `PATH` com `~/.local/bin` e `$(go env GOPATH)/bin`). Em terminal novo já configurado, pode pular.

2. **Gerar o código Go:**
   ```sh
   protoc --go_out=. --go-grpc_out=. proto/course_category.proto
   ```
   **O que faz:** Lê `proto/course_category.proto` e gera em `internal/pb/` as structs das mensagens (`--go_out`) e os stubs de cliente/servidor do `CategoryService` (`--go-grpc_out`).

### Arquivos gerados

Em `internal/pb/` são criados dois arquivos (**não edite** — são regenerados pelo `protoc`):

| Arquivo | Gerado por | Para que serve |
|---------|------------|-----------------|
| **`course_category.pb.go`** | `protoc-gen-go` | Structs das mensagens (`Category`, `CreateCategoryRequest`, `CategoryResponse`) com getters e serialização. |
| **`course_category_grpc.pb.go`** | `protoc-gen-go-grpc` | Interface e implementação do **cliente** (`CategoryServiceClient`, `CreateCategory`) e do **servidor** (`CategoryServiceServer`, `UnimplementedCategoryServiceServer`, `RegisterCategoryServiceServer`). |

No servidor você implementa `CategoryServiceServer` e chama `RegisterCategoryServiceServer`. No cliente você usa `NewCategoryServiceClient(conn)` e `CreateCategory`.

---

## O que é gRPC?

**gRPC** (gRPC Remote Procedure Call) é um framework de RPC de alto desempenho, desenvolvido pelo Google e aberto. Permite que um cliente chame métodos em um servidor em outra máquina como se fossem locais.

- **Baseado em contratos**: serviços e métodos definidos em arquivos `.proto`.
- **Protocol Buffers**: IDL e serialização binária por padrão.
- **Multiplataforma**: várias linguagens; cliente e servidor podem ser em linguagens diferentes.
- **HTTP/2**: transporte com multiplexação e streaming.

---

## RPC — Remote Procedure Call (Client → Server)

**RPC** é o modelo em que o **cliente** invoca um método que roda no **servidor** (outro processo/máquina), como uma chamada local.

1. Cliente envia requisição (método + parâmetros serializados).
2. Servidor executa o método e serializa o resultado.
3. Cliente recebe a resposta.

No gRPC, o cliente usa um *stub* gerado a partir do `.proto`.

```
[Cliente]  ---- requisição ---->  [Servidor]
[Cliente]  <---- resposta -------  [Servidor]
```

---

## Protocol Buffers

**Protocol Buffers** (protobuf) é o mecanismo de serialização e IDL usado por padrão no gRPC. Você define mensagens em `.proto` e o `protoc` gera código em várias linguagens.

- **IDL**: contrato entre cliente e servidor.
- **Binário**: payload compacto e rápido.
- **Tipagem**: campos com tipo definido; evolução compatível.

---

## gRPC vs Protocol Buffers

Não são alternativas: **Protocol Buffers** = formato e contrato dos dados; **gRPC** = framework de RPC que usa esse contrato para comunicação entre cliente e servidor. Em um projeto gRPC você escreve os `.proto` e o gRPC gera stubs e faz as chamadas remotas.

---

## Protocol Buffers vs JSON

| Aspecto | Protocol Buffers | JSON |
|--------|------------------|------|
| Formato | Binário | Texto |
| Tamanho / velocidade | Menor, mais rápido | Maior, mais lento |
| Schema | Obrigatório (`.proto`) | Opcional |
| Uso típico | Serviço a serviço, gRPC | APIs REST, front-end, debug |

---

## HTTP/2

O gRPC usa **HTTP/2**: multiplexação, compressão de cabeçalhos (HPACK), streaming bidirecional e uma única conexão para várias requisições.

---

## Tipos de API gRPC

| Tipo | Descrição |
|------|-----------|
| **Unary** | Uma requisição, uma resposta. |
| **Server streaming** | Cliente envia uma mensagem; servidor envia um stream de mensagens. |
| **Client streaming** | Cliente envia um stream; servidor devolve uma resposta. |
| **Bidirectional streaming** | Cliente e servidor enviam streams ao mesmo tempo. |

O método **CreateCategory** deste projeto é **Unary**.

---

## Em quais casos podemos utilizar?

gRPC é indicado para: comunicação entre microserviços, APIs de alta performance, streaming, ambientes com várias linguagens, mobile/IoT. Considere REST quando a API for voltada a browsers ou o ecossistema já for REST.

---

## REST vs gRPC

| Aspecto | REST | gRPC |
|--------|------|------|
| Modelo | Recursos + verbos HTTP | Métodos RPC (`.proto`) |
| Formato | Geralmente JSON | Protocol Buffers (binário) |
| Transporte | HTTP/1.1 comum | HTTP/2 |
| Streaming | Limitado | Nativo |

**Resumo:** REST para APIs públicas e front-end; gRPC para serviços internos, baixa latência e contratos fortes.

---

## Linguagens (suporte oficial)

gRPC suporta oficialmente: C++, C#/.NET, Dart, Go, Java, Kotlin, Node.js, Objective-C, PHP, Python, Ruby, Swift. Documentação por linguagem em [grpc.io/docs/languages](https://grpc.io/docs/languages/).

---

## Referências

- [Protocol Buffers](https://protobuf.dev/)
- [gRPC](https://grpc.io/)
- [Quick start — gRPC em Go](https://grpc.io/docs/languages/go/quickstart/)
- [Introdução ao gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Documentação gRPC](https://grpc.io/docs)
