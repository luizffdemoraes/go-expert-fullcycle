# gRPC

## O que é gRPC?

**gRPC** (gRPC Remote Procedure Call) é um framework de chamada de procedimento remoto (RPC) de alto desempenho, desenvolvido pelo Google e de código aberto. Ele permite que uma aplicação cliente chame métodos em uma aplicação servidor em outra máquina como se fossem objetos locais, facilitando a criação de aplicações e serviços distribuídos.

Principais características:

- **Baseado em contratos**: Você define um *service* em arquivos `.proto`, especificando os métodos que podem ser chamados remotamente, com parâmetros e tipos de retorno.
- **Protocol Buffers**: Por padrão, o gRPC usa **Protocol Buffers** (protobuf) como linguagem de definição de interface (IDL) e como formato de serialização das mensagens — o que resulta em payloads compactos e eficientes.
- **Multiplataforma**: Clientes e servidores gRPC podem rodar em diversos ambientes e serem escritos em qualquer uma das linguagens suportadas. Por exemplo, um servidor em Java pode atender clientes em Go, Python ou Ruby.
- **HTTP/2**: Utiliza HTTP/2 como transporte, permitindo multiplexação, compressão de cabeçalhos e comunicação bidirecional em um único canal.

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

## Referências

- [Introdução ao gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Linguagens suportadas](https://grpc.io/docs/languages/)
- [Documentação gRPC](https://grpc.io/docs)
