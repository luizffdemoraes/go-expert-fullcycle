# gRPC

## O que é gRPC?

**gRPC** (gRPC Remote Procedure Call) é um framework de chamada de procedimento remoto (RPC) de alto desempenho, desenvolvido pelo Google e de código aberto. Ele permite que uma aplicação cliente chame métodos em uma aplicação servidor em outra máquina como se fossem objetos locais, facilitando a criação de aplicações e serviços distribuídos.

Principais características:

- **Baseado em contratos**: Você define um *service* em arquivos `.proto`, especificando os métodos que podem ser chamados remotamente, com parâmetros e tipos de retorno.
- **Protocol Buffers**: Por padrão, o gRPC usa **Protocol Buffers** (protobuf) como linguagem de definição de interface (IDL) e como formato de serialização das mensagens — o que resulta em payloads compactos e eficientes.
- **Multiplataforma**: Clientes e servidores gRPC podem rodar em diversos ambientes e serem escritos em qualquer uma das linguagens suportadas. Por exemplo, um servidor em Java pode atender clientes em Go, Python ou Ruby.
- **HTTP/2**: Utiliza HTTP/2 como transporte, permitindo multiplexação, compressão de cabeçalhos e comunicação bidirecional em um único canal.

---

## Este projeto

Este repositório contém uma implementação em **Go** de um serviço gRPC para **categorias e cursos**: contrato da API em Protocol Buffers e camada de persistência em banco de dados, prontos para serem usados por um servidor e clientes gRPC.

### Para que serve

- **Expor uma API gRPC** de categorias: no contrato (`.proto`) está definido o `CategoryService` com o método **CreateCategory** (chamada unária: uma requisição, uma resposta). Outros métodos e o serviço de cursos podem ser estendidos a partir da mesma base.
- **Persistir dados**: a camada em `internal/database` implementa o acesso a tabelas **categories** e **courses** (PostgreSQL-compatível), com criação de categorias, listagem, e consultas por curso ou por categoria — essa camada será usada pelo handler gRPC para ler e gravar no banco.

### Como funciona

1. **Contrato (proto)** — Em `proto/course_category.proto` estão definidas as mensagens (`Category`, `CreateCategoryRequest`, `CategoryResponse`) e o serviço `CategoryService` com `CreateCategory`. Esse arquivo é a fonte da verdade para cliente e servidor; o código Go de stubs e mensagens é gerado com `protoc` (veja [Instalando compilador e plugins](#instalando-compilador-e-plugins)).
2. **Camada de dados** — Em `internal/database`:
   - **Category**: `Create`, `FindAll`, `FindByCourseID` — cria categoria (com UUID), lista todas e busca a categoria de um curso.
   - **Course**: `Create`, `FindAll`, `FindByCategoryID` — cria curso vinculado a uma categoria, lista todos e lista cursos por categoria.
3. **Fluxo (quando o servidor gRPC estiver implementado)** — O cliente chama `CreateCategory`; o servidor recebe a requisição, valida, chama `Category.Create` no banco e devolve `CategoryResponse` com a categoria criada.

### Motivação

- Praticar **gRPC em Go** com contrato primeiro (`.proto`) e código gerado.
- Separar bem **contrato da API** (proto), **regra/persistência** (database) e, futuramente, **transporte** (servidor gRPC).
- Ter uma base para evoluir para mais métodos (listar categorias, criar curso, etc.) e para outros tipos de chamada (streaming), mantendo a mesma estrutura de projeto.

---

## RPC — Remote Procedure Call (Client → Server)

**RPC** (Remote Procedure Call) é um modelo de comunicação em que um programa **cliente** invoca uma função ou método que está sendo executado em outro processo ou máquina — o **servidor** — como se fosse uma chamada local.

Fluxo resumido:

1. **Cliente**: envia uma requisição chamando um método remoto (com parâmetros serializados).
2. **Rede**: a mensagem trafega entre cliente e servidor.
3. **Servidor**: recebe a requisição, executa o método e serializa o resultado.
4. **Cliente**: recebe a resposta e segue a execução como em uma chamada local.

No gRPC, o cliente usa um *stub* (gerado a partir do `.proto`) que expõe os mesmos métodos do servidor, abstraindo a complexidade da rede e da serialização.

```
[Cliente]  ---- requisição (método + params) ---->  [Servidor]
[Cliente]  <---- resposta (resultado) ------------  [Servidor]
```

---

## Protocol Buffers

**Protocol Buffers** (protobuf) é o mecanismo de serialização estruturada usado por padrão no gRPC. Desenvolvido pelo Google, permite definir a estrutura dos dados em arquivos `.proto` e gerar código em várias linguagens para ler e escrever esses dados de forma eficiente.

- **IDL**: os arquivos `.proto` servem como contrato (interface) entre cliente e servidor.
- **Serialização binária**: os dados são codificados em formato binário compacto, menor e mais rápido que texto (ex.: JSON).
- **Tipagem**: campos têm tipo definido (string, int32, mensagens aninhadas, etc.), o que evita erros e facilita evolução da API com compatibilidade (ex.: campos opcionais, números de campo fixos).

Exemplo de definição:

```protobuf
message Person {
  string name = 1;
  int32 id = 2;
  bool active = 3;
}
```

O compilador `protoc` gera classes (ou structs) na linguagem escolhida para manipular essas mensagens.

---

## Protocol Buffers vs JSON

| Aspecto | Protocol Buffers | JSON |
|--------|------------------|------|
| **Formato** | Binário | Texto (legível) |
| **Tamanho** | Menor (payload mais compacto) | Maior (nomes de campos repetidos, sintaxe) |
| **Velocidade** | Serialização/deserialização geralmente mais rápida | Mais lenta e mais custo de CPU |
| **Legibilidade** | Não legível diretamente (precisa de ferramentas) | Legível em qualquer editor |
| **Schema** | Obrigatório (arquivo `.proto`) | Opcional (sem validação nativa de tipos) |
| **Uso típico** | Comunicação serviço-a-serviço, performance, gRPC | APIs REST, integração com front-end, debug manual |

**Resumo**: Protobuf é melhor para desempenho e eficiência entre serviços; JSON é mais conveniente para APIs voltadas a humanos e para inspeção manual (logs, Postman, etc.).

---

## gRPC vs Protocol Buffers

**gRPC** e **Protocol Buffers** não são alternativas: trabalham juntos. O Protocol Buffers cuida da forma dos dados e do contrato; o gRPC cuida da comunicação remota (RPC) e do transporte.

| Aspecto | Protocol Buffers | gRPC |
|--------|-------------------|------|
| **O que é** | Mecanismo de serialização estruturada e linguagem de definição (IDL) em arquivos `.proto`. | Framework de RPC (Remote Procedure Call) para chamar métodos entre cliente e servidor. |
| **Papel** | Define *como* os dados são estruturados e serializados (mensagens, tipos, campos). | Define *como* cliente e servidor se comunicam (métodos, streaming, transporte HTTP/2). |
| **Escopo** | Serialização + contrato dos dados. Pode ser usado sem gRPC (ex.: persistência, filas, outras APIs). | Chamadas remotas, streaming, load balancing, auth. Usa protobuf por padrão, mas pode usar outros formatos. |
| **Documentação** | [Protocol Buffers — protobuf.dev](https://protobuf.dev/) | [gRPC — grpc.io](https://grpc.io/) |

**Resumo**: Protocol Buffers = *formato e contrato dos dados*. gRPC = *framework de RPC* que usa esse contrato e esse formato para comunicação entre serviços. Em um projeto gRPC típico, você escreve os `.proto` (protobuf) e o gRPC usa esses arquivos para gerar stubs e fazer as chamadas remotas.

---

## REST vs gRPC

| Aspecto | REST | gRPC |
|--------|------|------|
| **Modelo** | Recursos (URLs) + verbos HTTP (GET, POST, PUT, DELETE) | Métodos RPC definidos no contrato (`.proto`) |
| **Formato** | Geralmente JSON (texto) | Protocol Buffers (binário) por padrão |
| **Transporte** | HTTP/1.1 (comum) ou HTTP/2 | HTTP/2 |
| **Contrato** | Documentação (OpenAPI, etc.) ou convenção | Schema obrigatório (`.proto`) com geração de código |
| **Streaming** | Limitado (SSE, chunked) | Nativo: server, client e bidirectional |
| **Performance** | Boa; overhead de texto e múltiplas requisições | Melhor: payload menor, multiplexação, uma conexão |
| **Browser** | Suporte nativo (fetch, XMLHttpRequest) | Requer gRPC-Web ou proxy |
| **Debug / ferramentas** | Fácil (curl, Postman, DevTools) | Requer ferramentas que entendam protobuf |
| **Adoção** | Muito usada em APIs públicas e front-end | Comum em back-end, microserviços, sistemas internos |

**Quando preferir REST:** APIs públicas, integração com front-end web, ecossistema já em REST, necessidade de debug simples e ferramentas universais.

**Quando preferir gRPC:** Comunicação entre serviços, baixa latência, alto throughput, streaming real, contratos fortes e código gerado em várias linguagens.

### Resumo comparativo: Texto/JSON (REST) vs Protocol Buffers (gRPC)

| Texto/JSON (REST) | Protocol Buffers (gRPC) |
|-------------------|-------------------------|
| Texto/JSON | Protocol Buffers |
| Undirecional | Bidirecional e assíncrono |
| Alta latência | Baixa latência |
| Sem contrato (maior chance de erros) | Contrato definido (`.proto`) |
| Sem suporte a streaming (Request/Response) | Suporte a streaming |
| Design pré-definido | Design é livre |
| Biblioteca de terceiros | Geração de código |

---

## HTTP/2

O gRPC usa **HTTP/2** como camada de transporte (não HTTP/1.1). Isso traz:

- **Multiplexação**: várias requisições/respostas no mesmo canal TCP, sem bloquear umas às outras.
- **Compressão de cabeçalhos (HPACK)**: menos overhead por requisição.
- **Streaming bidirecional**: cliente e servidor podem enviar múltiplas mensagens em sequência na mesma conexão.
- **Um único canal**: reduz número de conexões e latência em cenários com muitas chamadas.

Assim, o gRPC combina o modelo RPC + Protocol Buffers com as vantagens do HTTP/2 para comunicação eficiente e adequada a microserviços e streaming.

---

## Tipos de API gRPC

O gRPC suporta quatro tipos de chamada, definidos no `.proto` pela assinatura do método (uma requisição, uma resposta ou uso de `stream`).

### gRPC — API "Unary"

Uma requisição, uma resposta. É o modelo mais simples, equivalente a uma chamada de função remota clássica.

- **Cliente**: envia **uma** mensagem e espera **uma** resposta.
- **Servidor**: processa a mensagem e devolve **uma** resposta.

```protobuf
rpc GetUser (GetUserRequest) returns (GetUserResponse);
```

```
[Cliente]  ---- request ---->  [Servidor]
[Cliente]  <---- response ---  [Servidor]
```

---

### gRPC — API "Server streaming"

O cliente envia uma requisição e o servidor responde com um **fluxo** de mensagens.

- **Cliente**: envia **uma** mensagem.
- **Servidor**: envia **várias** mensagens em sequência (stream).

Útil para: listagens grandes, notificações em tempo real, envio de chunks de dados (ex.: arquivo, feed de eventos).

```protobuf
rpc ListItems (ListRequest) returns (stream Item);
```

```
[Cliente]  ---- request ---->  [Servidor]
[Cliente]  <---- msg 1 ------  [Servidor]
[Cliente]  <---- msg 2 ------  [Servidor]
[Cliente]  <---- msg N ------  [Servidor]
```

---

### gRPC — API "Client streaming"

O cliente envia um **fluxo** de mensagens; o servidor processa e devolve **uma** resposta (geralmente ao final).

- **Cliente**: envia **várias** mensagens em sequência.
- **Servidor**: envia **uma** resposta (tipicamente após receber e processar o stream).

Útil para: upload em chunks, envio de lotes de eventos ou métricas, agregação no servidor.

```protobuf
rpc UploadLogs (stream LogEntry) returns (UploadResult);
```

```
[Cliente]  ---- msg 1 ---->  [Servidor]
[Cliente]  ---- msg 2 ---->  [Servidor]
[Cliente]  ---- msg N ---->  [Servidor]
[Cliente]  <---- response -  [Servidor]
```

---

### gRPC — API "Bidirectional streaming"

Cliente e servidor enviam **fluxos** de mensagens ao mesmo tempo, de forma independente, na mesma conexão.

- **Cliente**: envia várias mensagens e pode receber várias ao longo do tempo.
- **Servidor**: envia várias mensagens e pode receber várias ao longo do tempo.
- A ordem e o momento de envio/recebimento podem ser independentes em cada direção.

Útil para: chat, jogos em tempo real, cooperação em tempo real, pipelines de processamento.

```protobuf
rpc Chat (stream ChatMessage) returns (stream ChatMessage);
```

```
[Cliente]  ---- msg ---->  [Servidor]
[Cliente]  <---- msg ----  [Servidor]
[Cliente]  ---- msg ---->  [Servidor]
[Cliente]  <---- msg ----  [Servidor]
         ... (ambos podem enviar quando quiserem)
```

---

## Em quais casos podemos utilizar?

O gRPC é indicado quando você precisa de:

| Cenário | Descrição |
|--------|-----------|
| **Comunicação entre microserviços** | Chamadas entre serviços internos com baixa latência e alto throughput. |
| **APIs de alta performance** | Quando desempenho e eficiência de rede são críticos (ex.: sistemas em tempo real, streaming). |
| **Streaming** | Comunicação em fluxo contínuo: **unário**, **stream do servidor**, **stream do cliente** ou **bidirecional**. |
| **Ambientes poliglotas** | Serviços escritos em linguagens diferentes que precisam se comunicar com um contrato bem definido (.proto). |
| **Mobile e IoT** | Uso de conexões eficientes e payloads menores, importantes para dispositivos com recursos limitados. |
| **Integração com APIs Google** | Muitas APIs do Google expõem interfaces gRPC, permitindo integrar funcionalidades do ecossistema Google. |

**Quando considerar alternativas (ex.: REST/JSON):**

- APIs voltadas principalmente para browsers (gRPC-Web existe, mas REST é mais comum).
- Quando a equipe ou o ecossistema já está fortemente baseado em REST e a mudança não traz ganhos claros.

---

## Linguagens (suporte oficial)

O gRPC possui suporte oficial para as seguintes linguagens e plataformas (conforme [grpc.io/docs/languages](https://grpc.io/docs/languages/)):

| Linguagem / Plataforma | Observações |
|------------------------|-------------|
| **C++** | Linux, macOS, Windows |
| **C# / .NET** | Windows, Linux, macOS (.NET Core 3.0+) |
| **Dart** | Windows, Linux, macOS (Dart 2.12+) — comum em Flutter |
| **Go** | Windows, Linux, macOS (Go 1.13+) |
| **Java** | Windows, Linux, macOS (Java 8+) |
| **Kotlin** | Windows, Linux, macOS (Kotlin 1.3+) |
| **Node.js** | Windows, Linux, macOS (Node v8+) |
| **Objective-C** | macOS, iOS |
| **PHP** | Linux, macOS (PHP 7.0+) |
| **Python** | Windows, Linux, macOS (Python 3.8+) |
| **Ruby** | Windows, Linux, macOS (Ruby 3.1+) |
| **Swift** | macOS, iOS |

Cada linguagem possui documentação com referência de API, tutoriais e guias de início rápido no site oficial do gRPC.

---

## Instalando compilador e plugins

Para gerar código gRPC a partir dos arquivos `.proto`, é necessário instalar o compilador **Protocol Buffers** (`protoc`) e os **plugins** da linguagem que você vai usar. Abaixo, um guia resumido para **Go** (o fluxo é análogo para outras linguagens).

### Pré-requisitos

- **Go** — uma das duas últimas versões principais ([Guia de instalação do Go](https://go.dev/doc/install)).
- **Compilador Protocol Buffers** — `protoc`, versão 3 ([Protoc Installation](https://protobuf.dev/programming-guides/protoc-installation/)).
- **Plugins do Go** para o compilador:
  1. Instale os plugins:
     ```sh
     go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
     go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
     ```
  2. Coloque o diretório dos binários do Go no `PATH` para o `protoc` encontrar os plugins:
     ```sh
     export PATH="$PATH:$(go env GOPATH)/bin"
     ```

### Gerando código a partir do `.proto`

No diretório do seu projeto (onde está o `.proto`), execute algo como:

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

Isso gera os arquivos de mensagens (ex.: `helloworld.pb.go`) e o código gRPC de cliente/servidor (ex.: `helloworld_grpc.pb.go`).

Para outros idiomas e detalhes completos (exemplo funcional, adicionar novos métodos, etc.), consulte o **[Quick start — gRPC em Go](https://grpc.io/docs/languages/go/quickstart/)**.

---

## Referências

- [Protocol Buffers — Documentação oficial](https://protobuf.dev/)
- [gRPC — Site oficial](https://grpc.io/)
- [Quick start — gRPC em Go](https://grpc.io/docs/languages/go/quickstart/)
- [Introdução ao gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Linguagens suportadas](https://grpc.io/docs/languages/)
- [Documentação gRPC](https://grpc.io/docs)
