# Solução para Erros ao Utilizar o Pacote go-sqlite3

## 1️⃣ Erro relacionado à variável CGO_ENABLED

#### **Mensagem de erro:**

```bash
[error] failed to initialize database, got error Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
panic: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub

goroutine 1 [running]:
main.main()
        C:/Users/lffde/go/src/github.com/fullcycle/curso-go/7-Apis/cmd/server/main.go:21 +0x186
exit status 2
```

### Causa:
Esse erro ocorre quando a variável de ambiente `CGO_ENABLED` está configurada como `0`. O pacote `go-sqlite3` exige que a variável `CGO_ENABLED` seja configurada como `1` para que a interface cgo funcione corretamente, já que a biblioteca depende de código nativo em C para acessar o banco de dados SQLite.

### Solução:
Para corrigir este erro, você precisa habilitar a variável `CGO_ENABLED` no seu terminal. Siga as instruções abaixo, dependendo do seu ambiente:

### Passos para corrigir o erro:
Habilitar o CGo

Antes de compilar e executar seu código, você precisa garantir que o CGo esteja habilitado. Em um terminal, execute:

### No Windows (PowerShell):

```bash
set CGO_ENABLED=1
```

### Instalar ferramentas necessárias para CGo

Como o go-sqlite3 usa CGo, você precisa de um compilador C no seu sistema. No Windows, você pode usar o MSYS2, que fornecem as ferramentas necessárias. 

### No Windows (MSYS2):

- Baixe e instale o MSYS2(https://www.msys2.org/).
- Instale as ferramentas C necessárias (como gcc e make) dentro do MSYS2.


1. **Atualizar o sistema e pacotes:**

```bash
pacman -Syu
```

2. **Instalar ferramentas de desenvolvimento básicas e a ferramenta de compilação MinGW:**

```bash
pacman -S base-devel mingw-w64-x86_64-toolchain
```

3. **Instalar o SQLite3 no MSYS2:**

```bash
pacman -S mingw-w64-x86_64-sqlite3
```

4. **Instalar o compilador GCC (GNU Compiler Collection) no MSYS2:**

```bash
pacman -S mingw-w64-x86_64-gcc
```

5. **Verificar a versão do GCC instalada:**

```bash
gcc --version
```

### Limpar o cache de módulos

Agora que você habilitou o CGo e instalou as ferramentas necessárias, limpe o cache de módulos e faça o download das dependências novamente:

```bash
go clean -modcache
```

#### No MSYS2:
Execute o comando abaixo para definir a variável `CGO_ENABLED` para `1`:

```bash
export CGO_ENABLED=1
```
#### No PowerShell (Windows):
Execute o comando abaixo:

```bash
$env:CGO_ENABLED=1
```

#### Verifique se a variável foi definida corretamente:
Você pode verificar o valor atual da variável CGO_ENABLED com o comando:

```bash
go env CGO_ENABLED
```

## 2️⃣ Erro relacionado ao GCC (Compilador C)

#### **Mensagem de erro:**
```bash
# runtime/cgo
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
```

### Causa:

Este erro ocorre porque o Go não consegue encontrar o compilador C gcc necessário para compilar o código nativo do SQLite ao usar o pacote go-sqlite3. O GCC não está instalado ou não está no PATH do sistema.

### Solução:

Para resolver esse erro, você precisa instalar o GCC e garantir que ele esteja acessível no PATH do seu sistema.

### Instalar o GCC no MSYS2:
Abra o terminal MSYS2 e execute o seguinte comando para instalar o GCC:

```bash
pacman -S mingw-w64-x86_64-gcc
```
Adicionar o GCC ao PATH no Windows:
Após a instalação, localize o diretório do GCC. Ele normalmente se encontra em:

```bash
C:\msys64\mingw64\bin
```

### Adicione esse diretório ao PATH do Windows:

- Pressione Win + X e selecione Sistema.
- Clique em Configurações avançadas do sistema.
- Na aba Avançado, clique em Variáveis de ambiente.
- Em Variáveis do sistema, edite a variável Path e adicione o diretório 

```bash
C:\msys64\mingw64\bin.
```
- Clique em OK para salvar as alterações.

### Verificar se o GCC está corretamente instalado:
Após adicionar o diretório ao PATH, feche e reabra o terminal, então execute o comando:

```bash
gcc --version
```
Se o GCC foi instalado corretamente, você verá a versão do compilador sendo exibida.

**OBS:** Para garantir que as alterações no PATH entrem em vigor, reinicie o terminal ou o prompt de comando.

[⬅ Voltar para o README principal](/README.MD)