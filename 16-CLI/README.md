# 16-CLI

CLI (Command Line Interface) em Go para gerenciamento de categorias e cursos, utilizando a biblioteca [Cobra](https://github.com/spf13/cobra).

## Introdução sobre CLI

Uma **CLI (Command Line Interface)** é uma interface de linha de comando que permite interagir com um programa através de comandos digitados no terminal. Em Go, uma das bibliotecas mais usadas para construir CLIs é o **Cobra**, que oferece:

- **Subcomandos** (ex.: `app create`, `app list`)
- **Flags** (opções como `--name`, `-v`)
- **Geração de ajuda** automática (`--help`)
- **Estrutura organizada** de comandos e subcomandos

O **cobra-cli** é um gerador de código que facilita a criação e a manutenção de projetos baseados em Cobra, gerando a estrutura de pastas e arquivos iniciais.

## Setup básico da aplicação

### Pré-requisitos

- [Go](https://go.dev/dl/) 1.19+ instalado
- `GOPATH` configurado (geralmente `~/go`) e `$GOPATH/bin` no `PATH` para usar binários instalados com `go install`

### Clonar e dependências

```bash
# Na pasta do projeto
cd 16-CLI

# Baixar dependências do módulo
go mod download
# ou
go mod tidy
```

O projeto usa:

- `github.com/google/uuid` para geração de IDs
- Pacotes em `internal/database` para entidades **Category** e **Course** (SQL)

## Inicializando projeto cobra-cli

Para criar e manter comandos com Cobra de forma prática, instale o **cobra-cli** e use-o para inicializar o projeto.

### 1. Instalar o cobra-cli

```bash
go install github.com/spf13/cobra-cli@latest
```

Confirme que o binário está no `PATH`:

```bash
cobra-cli --help
```

Se não for encontrado, adicione ao `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### 2. Inicializar o Cobra no projeto

Na raiz do módulo (`16-CLI`):

```bash
# Criar o comando raiz (root) e a pasta cmd
cobra-cli init
```

Isso cria (ou ajusta):

- `main.go` – ponto de entrada que chama o comando raiz
- `cmd/root.go` – definição do comando raiz e flags globais

### 3. Adicionar subcomandos

No exemplo deste projeto foram criados comandos simples de teste:

```bash
cobra-cli add ping
cobra-cli add pong
```

Após isso, ao executar:

```bash
go run main.go
```

Você verá algo semelhante a:

```text
Usage:
  16-CLI [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ping        A brief description of your command
  pong        A brief description of your command
```

### 4. Rodar a aplicação

Para apenas executar a aplicação durante o desenvolvimento:

```bash
go run main.go
```

Para gerar um binário:

```bash
go build -o cli && ./cli
```

Use `--help` no comando raiz ou em qualquer subcomando para ver a ajuda gerada pelo Cobra.

## Inicializando projeto cobra (passo a passo com terminal)

Esta seção descreve, passo a passo, os comandos executados e o que esperar no terminal com base na sessão real do projeto.

- **1. Ver comandos disponíveis no cobra-cli**

  ```bash
  cobra-cli
  ```

  **Esperado no terminal:** listagem de uso e comandos disponíveis (`add`, `completion`, `help`, `init`), indicando que o `cobra-cli` está instalado corretamente.

- **2. Organizar dependências do módulo**

  ```bash
  go mod tidy
  ```

  **Esperado no terminal:** em geral, nenhuma saída significativa quando tudo está correto; o comando apenas ajusta `go.mod` e `go.sum`.

- **3. Inicializar a aplicação Cobra**

  ```bash
  cobra-cli init
  ```

  **Esperado no terminal:** mensagem parecida com:

  ```text
  Your Cobra application is ready at
  /home/lffm1994/workspace/go-expert-fullcycle/16-CLI
  ```

- **4. Primeiro `go run main.go`**

  ```bash
  go run main.go
  ```

  **Esperado no terminal:** a descrição padrão gerada pelo Cobra, algo como:

  ```text
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
  ```

- **5. Adicionar comando `ping`**

  ```bash
  cobra-cli add ping
  ```

  **Esperado no terminal:** mensagem indicando criação do comando, por exemplo:

  ```text
  ping created at /home/lffm1994/workspace/go-expert-fullcycle/16-CLI
  ```

  Ao rodar novamente:

  ```bash
  go run main.go
  ```

  você verá o comando `ping` listado em **Available Commands**.

- **6. Adicionar comando `pong`**

  ```bash
  cobra-cli add pong
  ```

  **Esperado no terminal:**

  ```text
  pong created at /home/lffm1994/workspace/go-expert-fullcycle/16-CLI
  ```

  Rodando mais uma vez:

  ```bash
  go run main.go
  ```

  agora tanto `ping` quanto `pong` aparecem na lista de comandos disponíveis.

- **7. Geração de autocomplete para zsh**

  A forma correta é rodar com o arquivo `main.go`:

  ```bash
  go run main.go completion zsh
  ```

  **Importante:** o comando incorreto

  ```bash
  go run main completion zsh
  ```

  resulta em erro:

  ```text
  package main is not in std (/usr/local/go/src/main)
  ```

  Usando `main.go` corretamente, o esperado é a saída de um script grande de completion para zsh (iniciando com `#compdef 16-CLI`), que você pode redirecionar para um arquivo e carregar no seu shell.

## Criando nossos primeiros comandos

Nesta etapa, criamos um comando chamado `teste` em `cmd/teste.go`, com uma flag `--comando` (atalho `-c`) que aceita os valores `ping` ou `pong`. O comando imprime no terminal o valor escolhido ou informa que o comando é inválido.

- **Definição do comando `teste`**

  O comando é registrado no Cobra com:

  - `Use: "teste"` – nome do comando (`16-CLI teste`)
  - Flag `--comando` / `-c` – texto de ajuda: `Escolha ping ou pong`
  - A flag é obrigatória (`MarkFlagRequired("comando")`)

- **1. Executando com `--comando=ping`**

  ```bash
  go run main.go teste --comando=ping
  ```

  **Esperado no terminal:**

  ```text
  ping
  ```

- **2. Executando com `-c ping` (forma abreviada)**

  ```bash
  go run main.go teste -c ping
  ```

  **Esperado no terminal:**

  ```text
  ping
  ```

- **3. Tentando usar uma flag escrita de forma incorreta**

  ```bash
  go run main.go teste --comand ping
  ```

  **Esperado no terminal:** erro de flag desconhecida e ajuda do comando:

  ```text
  Error: unknown flag: --comand
  Usage:
    16-CLI teste [flags]

  Flags:
    -c, --comando string   Escolha ping ou pong
    -h, --help             help for teste

  exit status 1
  ```

  Esse exemplo mostra como o Cobra valida flags automaticamente e exibe a ajuda correta quando algo está errado.

- **4. Flag curta escrita de forma incorreta**

  ```bash
  go run main.go teste -comand ping
  ```

  Como o código do comando trata qualquer valor diferente de `ping` ou `pong` como inválido, o resultado é:

  ```text
  comando inválido
  ```

  Aqui, mesmo que a CLI rode, o valor de `comando` não é reconhecido na lógica do `switch`, e o comportamento padrão é avisar que o comando é inválido.


---

**Resumo:** instale o gerador com `go install github.com/spf13/cobra-cli@latest`, use `cobra-cli init` para inicializar o projeto e `cobra-cli add <nome>` para criar novos comandos.
