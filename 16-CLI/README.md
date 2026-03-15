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

## Comandos encadeados

No Cobra, **comandos encadeados** (ou hierárquicos) são subcomandos anexados a outro comando, formando cadeias como `16-CLI category create` e `16-CLI category list`. Isso organiza a CLI por contexto (por exemplo, tudo relacionado a categorias fica sob `category`).

### Estrutura implementada

- **Raiz:** `16-CLI` (`cmd/root.go`)
- **Comando intermediário:** `category` (`cmd/category.go`) — anexado à raiz com `rootCmd.AddCommand(categoryCmd)`.
- **Subcomandos de category:**
  - `create` (`cmd/create.go`) — anexado com `categoryCmd.AddCommand(createCmd)`.
  - `list` (`cmd/list.go`) — anexado com `categoryCmd.AddCommand(listCmd)`.

Assim, a árvore fica:

```text
16-CLI
└── category
    ├── create
    └── list
```

### Como usar

| Comando | Descrição |
|--------|-----------|
| `go run main.go category` | Exibe a ajuda do comando `category` (e lista os subcomandos `create` e `list`). |
| `go run main.go category create` | Executa a ação de criação (atualmente imprime `create called`). |
| `go run main.go category list` | Executa a ação de listagem (atualmente imprime `list called`). |

### Exemplos no terminal

**Ajuda do comando `category`:**

```bash
go run main.go category --help
```

**Esperado:** saída com uso `16-CLI category [command]` e lista de subcomandos disponíveis (`create`, `list`).

**Chamar create e list:**

```bash
go run main.go category create
# create called

go run main.go category list
# list called
```

### Comandos no terminal para criar a estrutura encadeada

Use a flag `-p` (ou `--parent`) do cobra-cli para indicar o comando pai. O gerador já cria o arquivo com `parentCmd.AddCommand(novoCmd)` no `init()`.

**Criar o comando pai (nível raiz):**

```bash
cobra-cli add category
```

**Esperado:** criação de `cmd/category.go` e registro em `rootCmd` (ou você anexa manualmente em `root.go` se precisar).

**Criar subcomandos já vinculados ao `category`:**

```bash
cobra-cli add create -p 'categoryCmd'
cobra-cli add list -p 'categoryCmd'
```

**Esperado no terminal:** mensagens como:

```text
create created at /home/lffm1994/workspace/go-expert-fullcycle/16-CLI
list created at /home/lffm1994/workspace/go-expert-fullcycle/16-CLI
```

Os arquivos `cmd/create.go` e `cmd/list.go` são gerados com `categoryCmd.AddCommand(createCmd)` e `categoryCmd.AddCommand(listCmd)` no `init()`, então não é necessário alterar manualmente o pai do comando.

### Como criar subcomandos encadeados (resumo)

1. Gerar o comando pai (ex.: `category`): `cobra-cli add category`.
2. Gerar cada subcomando passando o pai com `-p 'categoryCmd'`:
   - `cobra-cli add create -p 'categoryCmd'`
   - `cobra-cli add list -p 'categoryCmd'`

Assim, `create` e `list` ficam sob `category`, e a CLI oferece os comandos encadeados `16-CLI category create` e `16-CLI category list`.

## Flags locais vs globais

No Cobra, **flags locais** valem só para aquele comando; **flags globais** (persistentes) valem para o comando e para todos os seus subcomandos. A escolha entre um e outro define onde a opção pode ser usada na linha de comando.

### Diferenças

| Aspecto | Flag local (`Flags()`) | Flag global / persistente (`PersistentFlags()`) |
|--------|-------------------------|--------------------------------------------------|
| **Escopo** | Apenas no comando em que foi definida | No comando e em todos os subcomandos (e sub-subcomandos) |
| **Uso** | `comando --flag` | `comando --flag` ou `comando subcomando --flag` |
| **Definição** | `cmd.Flags().String("name", "", "desc")` | `cmd.PersistentFlags().String("name", "", "desc")` |
| **Quando usar** | Opção específica daquela ação | Opção comum a várias ações (ex.: `--name` em create e list) |

### Alterações atuais no projeto

**1. Comando raiz (`cmd/root.go`) — flag local**

A flag `toggle` é **local**: só existe quando se chama o comando raiz, não aparece em `category`, `create` ou `list`.

```go
rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
```

**Uso no terminal:**

```bash
go run main.go -t
go run main.go --toggle
# Válido; a flag é reconhecida.

go run main.go category -t
# Inválido; -t não existe no comando category.
```

**2. Comando `category` (`cmd/category.go`) — flag persistente**

A flag `name` é **global (persistente)** no comando `category`: ela é herdada por `create` e `list`.

```go
categoryCmd.PersistentFlags().String("name", "", "Name of the category")   // flag global
// categoryCmd.Flags().String("name", "", "Name of the category")          // flag local (comentada)
```

Com `PersistentFlags()`, tanto `category` quanto `category create` e `category list` aceitam `--name`:

**Uso no terminal:**

```bash
go run main.go category --name "Minha Categoria"
go run main.go category create --name "Nova"
go run main.go category list --name "Filtro"
# Todos válidos; --name está disponível em category e nos subcomandos.
```

Se a flag fosse **local** (`categoryCmd.Flags().String("name", ...)`), apenas `go run main.go category --name "..."` funcionaria; `category create --name` e `category list --name` não reconheceriam `--name`.

### Resumo prático

- **Local:** `cmd.Flags()` — só naquele comando.
- **Global (persistente):** `cmd.PersistentFlags()` — naquele comando e em todos os subcomandos.

No projeto, a raiz usa flag local (`toggle`) e o comando `category` usa flag persistente (`name`) para que `create` e `list` possam receber `--name` sem redefinir a flag em cada um.

## Manipulando flags

Este tópico descreve como as flags são definidas, lidas no código e usadas no terminal, com base nas implementações do comando `category`.

### Implementações no comando `category` (`cmd/category.go`)

O comando `category` define três flags **persistentes** (disponíveis também para `category create` e `category list`):

| Flag (longa) | Atalho | Tipo   | Valor padrão | Descrição                    |
|--------------|--------|--------|--------------|-----------------------------|
| `--name`     | `-n`   | string | `"Y"`        | Nome da categoria           |
| `--exists`   | `-e`   | bool   | `false`      | Indica se a categoria existe |
| `--id`       | `-i`   | int    | `0`          | ID da categoria             |

**Definição no código:**

A flag `name` está ligada a uma variável por referência (`StringVarP`); as demais usam as formas que retornam valor e são lidas no `Run` com `Get*`:

```go
categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if the category exists")
categoryCmd.PersistentFlags().IntP("id", "i", 0, "ID of the category")
```

- `StringP`, `BoolP`, `IntP`: o sufixo **P** indica que a flag tem forma **curta** (segundo argumento: `"n"`, `"e"`, `"i"`).
- `StringVarP(&var, ...)`: o valor da flag é escrito na variável `var` por referência; veja o tópico **Flags mudando valor por referência**.

**Leitura das flags no `Run`:**

No handler do comando, os valores são obtidos com `GetString`, `GetBool` e `GetInt`:

```go
Run: func(cmd *cobra.Command, args []string) {
    name, _ := cmd.Flags().GetString("name")
    fmt.Println("Category called with name:", name)
    exists, _ := cmd.Flags().GetBool("exists")
    fmt.Println("Category exists:", exists)
    id, _ := cmd.Flags().GetInt("id")
    fmt.Println("Category ID:", id)
},
```

Sempre use o **nome longo** da flag (`"name"`, `"exists"`, `"id"`) em `GetString`/`GetBool`/`GetInt`, independentemente de o usuário ter passado `-n` ou `--name` no terminal.

### Comando executado no terminal

Exemplo combinando as três flags (forma curta com `=`, forma longa com `=`, e flag booleana):

```bash
go run main.go category -n=categoria -e --id=10
```

**Significado:**

- `-n=categoria` — define o nome da categoria como `categoria`.
- `-e` — ativa a flag booleana `exists` (equivalente a `--exists`).
- `--id=10` — define o ID como `10`.

**Saída esperada:**

```text
Category called with name: categoria
Category exists: true
Category ID: 10
```

### Formas válidas de passar flags

O Cobra aceita várias sintaxes; abaixo exemplos equivalentes para o mesmo resultado.

**String (`name`):**

```bash
go run main.go category -n=categoria
go run main.go category -n categoria
go run main.go category --name=categoria
go run main.go category --name categoria
```

**Bool (`exists`):**

```bash
go run main.go category -e
go run main.go category --exists
```

**Int (`id`):**

```bash
go run main.go category -i=10
go run main.go category -i 10
go run main.go category --id=10
go run main.go category --id 10
```

Como as flags de `category` são persistentes, os subcomandos também aceitam as mesmas flags:

```bash
go run main.go category create -n "Nova" -e --id=5
go run main.go category list --name "Filtro"
```

## Flags mudando valor por referência

Em vez de só ler o valor da flag no `Run` com `GetString`, `GetBool`, `GetInt`, você pode **atrelar a flag a uma variável**: o Cobra preenche essa variável por referência ao fazer o parse da linha de comando. Assim, no `Run` você usa a variável diretamente, sem chamar `Get*`.

### Implementação no projeto

No comando `category` (`cmd/category.go`), a flag `name` está ligada à variável de pacote `category` por referência.

**1. Declarar a variável (escopo de pacote):**

```go
var category string
```

**2. Definir a flag com a variante `*Var` / `*VarP`:**

```go
categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
```

- `StringVarP` recebe um **ponteiro** (`&category`) como primeiro argumento.
- Quando o usuário passa `-n=categoria` ou `--name=categoria`, o Cobra atribui o valor à variável `category` antes de executar o `Run`.
- No `Run`, você pode usar `category` diretamente ou continuar usando `cmd.Flags().GetString("name")` — ambos refletem o valor informado na linha de comando.

**3. Uso no `Run` (opcional):**

Se quiser usar a variável em vez de `GetString`:

```go
Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Category called with name:", category)  // variável já preenchida
    exists, _ := cmd.Flags().GetBool("exists")
    fmt.Println("Category exists:", exists)
    id, _ := cmd.Flags().GetInt("id")
    fmt.Println("Category ID:", id)
},
```

### Comando executado no terminal

```bash
go run main.go category -n=categoria -e --id=10
```

Com a flag `name` atrelada a `&category`, o Cobra faz `category = "categoria"` antes do `Run`. A saída esperada continua:

```text
Category called with name: categoria
Category exists: true
Category ID: 10
```

### Resumo: definição com valor vs por referência

| Abordagem | Definição | Leitura |
|-----------|-----------|---------|
| **Só valor** | `Flags().StringP("name", "n", "Y", "desc")` | `cmd.Flags().GetString("name")` no `Run` |
| **Por referência** | `Flags().StringVarP(&category, "name", "n", "Y", "desc")` | Usar a variável `category` no `Run` (já preenchida) |

As variantes `*Var` / `*VarP` existem para todos os tipos comuns: `StringVarP`, `BoolVarP`, `IntVarP`, `Int64VarP`, etc. Usar referência é útil quando você quer reutilizar o valor em vários lugares ou em funções que recebem a variável por parâmetro.

## Entendendo hooks

No Cobra, **hooks** são funções que rodam em momentos específicos do ciclo de vida do comando: antes do comando principal (`PreRun`), depois (`PostRun`) ou em vez dele com retorno de erro (`RunE`). Eles servem para inicializar recursos, validar argumentos, fazer log ou limpeza.

### Ordem de execução

Para um comando (ex.: `category`), a ordem típica é:

1. **PersistentPreRun** (do comando e dos pais, da raiz para o comando atual)
2. **PreRun** (apenas do comando atual)
3. **Run** ou **RunE** (ação principal; se **RunE** estiver definido, ele é usado no lugar de **Run**)
4. **PostRun** (apenas do comando atual)
5. **PersistentPostRun** (do comando e dos pais, do comando atual até a raiz)

### Implementação no projeto (`cmd/category.go`)

O comando `category` define três hooks para ilustrar o fluxo:

**PreRun — executado antes do `Run`/`RunE`:**

```go
PreRun: func(cmd *cobra.Command, args []string) {
    fmt.Println("Chamado antes da execução do Run")
},
```

**PostRun — executado depois do `Run`/`RunE`:**

```go
PostRun: func(cmd *cobra.Command, args []string) {
    fmt.Println("Chamado depois da execução do Run")
},
```

**RunE — ação principal com retorno de erro:**

Quando definido, o Cobra usa **RunE** em vez de **Run**. Se a função retornar um erro não nulo, a execução é interrompida e o programa pode sair com código de erro.

```go
RunE: func(cmd *cobra.Command, args []string) error {
    return fmt.Errorf("Ocorreu um erro")
},
```

Se **Run** e **RunE** estiverem ambos definidos no mesmo comando, apenas **RunE** é executado. No exemplo acima, como `RunE` retorna erro, a saída do programa será o erro e o código de saída 1; **PostRun** ainda é chamado após o **RunE**.

### Resumo dos hooks

| Hook | Escopo | Uso típico |
|------|--------|------------|
| **PersistentPreRun** | Comando + todos os subcomandos | Abrir conexão com DB, carregar config global |
| **PreRun** | Só o comando atual | Validar args/flags antes da ação |
| **Run** | Ação principal (sem retorno de erro) | Lógica do comando |
| **RunE** | Ação principal (retorna `error`) | Mesmo que Run, mas com tratamento de erro integrado |
| **PostRun** | Só o comando atual | Log, métricas, limpeza local |
| **PersistentPostRun** | Comando + todos os subcomandos | Fechar conexões, flush de buffers |

### Exemplo de saída no terminal

Com a implementação atual, ao rodar:

```bash
go run main.go category -n=categoria -e --id=10
```

a ordem da saída será: primeiro **PreRun** (“Chamado antes da execução do Run”), depois **RunE** (que retorna erro e pode ser exibido), em seguida **PostRun** (“Chamado depois da execução do Run”). O **Run** não é executado quando **RunE** está definido.

Para ver **PreRun** e **PostRun** junto com a lógica do **Run**, remova ou comente o **RunE** no `categoryCmd`; assim o Cobra usará o **Run** e você verá as três mensagens na ordem: PreRun → Run → PostRun.

---

**Resumo:** instale o gerador com `go install github.com/spf13/cobra-cli@latest`, use `cobra-cli init` para inicializar o projeto e `cobra-cli add <nome>` para criar novos comandos.
