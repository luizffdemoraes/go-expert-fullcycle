# 19-DI — Dependency Injection

Exemplo em Go de composição de dependências com **Google Wire** (geração de código em tempo de compilação) e **injeção direta** no domínio da aplicação.

## O que é injeção direta neste projeto

**Injeção direta** aqui significa que cada dependência é passada de forma **explícita** pelos construtores, sem localizador de serviços, singletons globais ou reflexão em tempo de execução para montar o grafo de objetos.

1. **No pacote `product`**  
   - `NewProductRepository(db *sql.DB)` recebe o banco diretamente.  
   - `NewProductUseCase(repository ProductRepositoryInterface)` recebe a **interface** do repositório, não o tipo concreto — o caso de uso depende de um contrato, facilitando testes e troca de implementação.

2. **No ponto de composição (`main`)**  
   - Quem “liga” tudo é `NewUseCase(db)`, que devolve um `*product.ProductUseCase` já configurado.  
   - A função gerada em `wire_gen.go` faz exatamente o que você faria na mão: criar o repositório com `db`, passar o repositório ao caso de uso e retornar o caso de uso — **cadeia explícita de `New...`**, sem mágica em runtime.

Ou seja: a **regra de negócio** não procura dependências; quem monta o grafo (`main` → `NewUseCase`) **injeta** tudo de forma direta e visível.

## Papel do Wire

O [Wire](https://github.com/google/wire) não é um container IoC em runtime. Ele **analisa** `wire.go` (com build tag `wireinject`) e **gera** o código em `wire_gen.go` com as chamadas aos construtores na ordem certa.

- **`wire.go`**: descreve o grafo (`wire.Build`, `wire.NewSet`, `wire.Bind` da interface ao tipo concreto).  
- **`wire_gen.go`**: código gerado — não editar manualmente; contém a `NewUseCase` com a montagem explícita.

O import do driver SQLite com `_` em `wire_gen.go` existe para que o pacote seja registrado quando o grafo é compilado junto ao binário (mesmo padrão do `main.go`).

### Como usar o Wire neste projeto

1. **Build tags (`wireinject`)**  
   - `wire.go` começa com `//go:build wireinject` — esse arquivo **só** é compilado quando você roda a ferramenta `wire` (ela usa essa tag para analisar o injetor).  
   - `wire_gen.go` usa `//go:build !wireinject` — é o que o `go build` / `go run` usam no dia a dia. Assim o binário não mistura o “rascunho” do injetor com o código gerado.

2. **Definir o injetor em `wire.go`**  
   - Crie uma função com a **assinatura desejada** (ex.: `NewUseCase(db *sql.DB) *product.ProductUseCase`).  
   - No corpo, chame `wire.Build(...)` listando tudo que o Wire precisa para montar o retorno: construtores (`product.NewProductUseCase`) e, opcionalmente, `wire.NewSet(...)` com `wire.Bind` quando um parâmetro é **interface** e a implementação é um **tipo concreto** (neste repo: `ProductRepositoryInterface` → `*ProductRepository`).  
   - O `return` no final é só um valor “placeholder”; o Wire ignora e substitui pela geração.

3. **Gerar o código**  
   - Na raiz do módulo: `go generate .` (o `//go:generate` em `wire_gen.go` invoca o comando `wire`).  
   - Depois disso, confira se `wire_gen.go` foi atualizado e rode `go build` ou `go run .`.

4. **Alterar dependências**  
   - Ao incluir um novo tipo no grafo (novo `New...` ou novo binding), **edite só `wire.go`** e rode `go generate .` de novo.

5. **Ferramenta**  
   - O `go generate` usa o módulo `github.com/google/wire`; ter `github.com/google/wire` no `go.mod` (como neste projeto) permite o `go run` do gerador sem instalar o binário `wire` no PATH — opcionalmente você pode usar `go install github.com/google/wire/cmd/wire@latest` e rodar `wire` na pasta do pacote.

## Bibliotecas de DI e sua utilização

Em Go é comum combinar **construtores explícitos** (`New...`) com uma ferramenta que monta o grafo. Duas abordagens frequentes:

### [Uber Fx](https://github.com/uber-go/fx)

O [Fx](https://github.com/uber-go/fx) (`go.uber.org/fx`) é um **framework de aplicação** com injeção de dependências em **tempo de execução**. Você declara **provedores** (`fx.Provide`) e **pontos de entrada** (`fx.Invoke`), e o Fx resolve tipos, ordem de construção e **ciclo de vida** (por exemplo, hooks de inicialização e encerramento).

**Uso típico:** serviços de longa duração (HTTP/gRPC, workers), módulos compartilháveis entre equipes, necessidade de startup/shutdown coordenado e substituição de implementações em testes com `fx.Replace` / `fx.Decorate`.

**Trade-off:** há um container e reflexão em runtime; o binário inclui o framework; a curva de aprendizado é um pouco maior, mas a documentação e o ecossistema são maduros.

### [Google Wire](https://github.com/google/wire)

O [Wire](https://github.com/google/wire) gera **código Go comum** em tempo de compilação: o grafo vira uma função com `:=` e chamadas a `New...`, **sem container em runtime** e sem reflexão para montar dependências.

**Uso típico:** quem prefere **código gerado explícito** fácil de ler e depurar, binários mínimos e comportamento previsível no `go build`.

**Observação:** o repositório oficial foi arquivado e o projeto não é mais mantido ativamente; mesmo assim, a ferramenta continua utilizada e o código gerado permanece válido — avalie forks ou alternativas (por exemplo Fx) para projetos novos de longo prazo.

### Comparando rápido

| | **Fx** | **Wire** |
|---|--------|----------|
| Momento da montagem | Runtime (ao subir a app) | Compile time (`go generate`) |
| Estilo | API declarativa (`Provide` / `Invoke`) | Descrição do grafo + codegen |
| Adequado quando | App com lifecycle, muitos serviços | Grafo estável, preferência por Go “puro” gerado |

### Outras bibliotecas parecidas

Nada substitui Fx ou Wire de forma idêntica, mas estas cobrem o **mesmo tipo de problema** (montar o grafo de dependências):

**Estilo runtime (container na subida da app), no mesmo campo do Fx**

- **[Uber Dig](https://github.com/uber-go/dig)** (`go.uber.org/dig`) — Toolkit de DI com reflexão na **inicialização**. O próprio **Fx** é implementado em cima do Dig: se quiser controle fino sem o framework completo, o Dig é a camada base.
- **[samber/do](https://github.com/samber/do)** — Container com **generics** (Go 1.18+), registro por tipo/nome, escopos e checagens úteis; outra opção de DI em runtime com foco em type-safety.
- **[Ore](https://github.com/firasdarwish/ore)** — Container mais leve, com lifetimes no estilo singleton/scoped/transient (inspirado em outros ecossistemas). Menos adotado que Fx/Dig, mas com ideia semelhante.

**Estilo compile time (código gerado), no mesmo campo do Wire**

- Em Go há **poucas** opções maduras além do Wire: muitas equipes usam **wiring manual** no `main`/cmd ou ferramentas de codegen **case a case**. Vale pesquisar projetos mais novos no mesmo nicho e avaliar manutenção e compatibilidade com seu `go.mod` antes de adotar.

**Este repositório** usa **Wire** para gerar o injetor; o domínio (`product`) continua agnóstico: só precisa de construtores e interfaces, o que também facilitaria migrar para Fx, Dig, outro container ou montagem manual no `main` se quiser.

## Como executar

```bash
go run .
```

## Regenerar o código do Wire

Após alterar `wire.go`, rode `go generate .` (detalhes no tópico [Como usar o Wire neste projeto](#como-usar-o-wire-neste-projeto)). A dependência `github.com/google/wire` no `go.mod` já cobre o comando do `go generate`.

## Estrutura

| Caminho | Função |
|--------|--------|
| `main.go` | Abre o SQLite, chama `NewUseCase(db)` e usa o caso de uso |
| `wire.go` | Definição do injetor (apenas com tag `wireinject`) |
| `wire_gen.go` | Injetor gerado (tag `!wireinject`) |
| `product/` | Entidade, repositório e caso de uso com injeção por construtor |
