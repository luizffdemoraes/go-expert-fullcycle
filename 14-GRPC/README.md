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

### Criando servidor gRPC

O ponto de entrada do servidor está em **`internal/cmd/grpcServer/main.go`**. Ele sobe o gRPC na porta **50051** e registra o `CategoryService` com reflexão habilitada (útil para ferramentas como `grpcurl`/evans).

**Ordem da implementação:**

1. **Conexão com o banco** — `sql.Open("sqlite3", "db.sqlite")` (ou troque para PostgreSQL conforme a camada em `internal/database`). O `defer db.Close()` garante o fechamento ao encerrar o processo.

2. **Camada de dados** — Cria-se o repositório de categorias: `categoryDB := database.NewCategory(db)`.

3. **Camada de serviço** — Cria-se o serviço que implementa o gRPC: `categoryService := service.NewCategoryService(*categoryDB)`. Esse é o tipo que implementa `CategoryServiceServer` e será registrado no servidor.

4. **Servidor gRPC** — `grpc.NewServer()` cria o servidor; `pb.RegisterCategoryServiceServer(grpcServer, categoryService)` registra o `CategoryService`; `reflection.Register(grpcServer)` ativa a reflexão para descoberta de serviços.

5. **Listen e Serve** — `net.Listen("tcp", ":50051")` abre a porta; `grpcServer.Serve(lis)` bloqueia e atende as chamadas RPC (por exemplo, `CreateCategory`).

**Como rodar:** a partir da raiz do projeto, execute o pacote do servidor (ex.: `go run ./internal/cmd/grpcServer` ou, se o módulo estiver em `cmd/grpcServer`, `go run .` de dentro desse diretório). O servidor fica ouvindo em `localhost:50051` até ser interrompido.

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
