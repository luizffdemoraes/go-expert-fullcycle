**Comandos de Compilação de Projetos.:**

## 1. go build nome-do-arquivo.go
O comando go build nome-do-arquivo.go é usado para compilar um arquivo Go específico e gerar um executável a partir do código-fonte.

🔹 Exemplo de uso
```bash
go build main.go
```

🔹 Como funciona?
O Go compila o código-fonte do arquivo nome-do-arquivo.go.
Se não houver erros, ele gera um executável no mesmo diretório.
Windows: nome-do-arquivo.exe
Linux/macOS: nome-do-arquivo

## 2. GOOS no Go
GOOS é uma variável de ambiente no Go que define o sistema operacional para o qual o código será compilado. Ele permite a compilação cruzada, ou seja, gerar um binário para um sistema diferente do atual.

## 2.1 Comando para listar combinações para GOOS
O comando go tool dist list lista todas as combinações possíveis de GOOS (sistemas operacionais) e GOARCH (arquiteturas) suportadas pelo compilador Go para compilação cruzada.

🔹 Exemplo de uso
```bash
go tool dist list
```

## 3. O que é GOARCH no Go?
GOARCH é uma variável de ambiente no Go que define a arquitetura da CPU para a qual o código será compilado. Ela permite a compilação cruzada ao ser combinada com GOOS para gerar binários para diferentes plataformas e arquiteturas.

## 3.1 Exibir o sistema operacional e arquitetura

O comando go env GOOS GOARCH exibe o sistema operacional (GOOS) e a arquitetura (GOARCH) do ambiente atual em que o Go está rodando.

🔹 Exemplo de uso
```bash
go env GOOS GOARCH
```

Exemplo de Uso
Se estiver no Windows e quiser compilar um binário para Linux, você pode usar:

```bash
GOOS=linux GOARCH=amd64 go build -o meu_programa_linux
```

Você pode verificar quais valores são suportados pelo seu Go com:

```bash
go env GOOS
```

### 🔹 Valores Comuns para GOOS

| GOOS     | Descrição |
|----------|----------|
| `linux`  | Linux    |
| `windows`| Windows  |
| `darwin` | macOS    |
| `freebsd`| FreeBSD  |


### 🔹 Valores Comuns para GOARCH

| GOARCH   | Arquitetura       | Descrição                                         |
|----------|------------------|---------------------------------------------------|
| `amd64`  | x86_64 (64 bits) | Arquitetura mais comum em PCs                    |
| `386`    | x86 (32 bits)    | CPUs mais antigas (32 bits)                      |
| `arm`    | ARM de 32 bits   | Usado em dispositivos móveis                     |
| `arm64`  | ARM de 64 bits   | Usado em novos chips Apple e servidores ARM      |
| `riscv64`| RISC-V 64 bits   | Arquitetura emergente                            |
| `wasm`   | WebAssembly      | Para rodar no navegador                          |



[⬅ Voltar para o README principal](/README.MD)
