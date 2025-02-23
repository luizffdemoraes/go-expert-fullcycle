**Comandos.:**

1️⃣ go run .

🔹 O que faz?

Compila e executa todos os arquivos Go do diretório atual que pertencem ao mesmo pacote.
Diferente de go run main.go, que executa apenas um arquivo específico.

2️⃣ go mod init <nome-do-modulo>

🔹 O que faz?

Cria um arquivo go.mod, inicializando um módulo Go no projeto.
Define o namespace do seu código para permitir importações de pacotes locais e externos.

🔹 Exemplo de uso
```bash
go mod init github.com/usuario/meu-projeto
```

3️⃣ go mod tidy

🔹 O que faz?

Remove dependências não utilizadas do go.mod e go.sum.
Baixa dependências faltantes no projeto.


4️⃣ go mod edit -replace github.com/fullcycle/curso-go/5-Packaging/3/math=../math

🔹 O que faz?

Esse comando modifica o arquivo go.mod para substituir um módulo remoto (github.com/fullcycle/curso-go/5-Packaging/3/math) por um caminho local (../math).

✅ Isso é útil para testar mudanças localmente sem precisar publicar o módulo.


5️⃣ go work init ./math ./sistema  

🔹 O que faz?

Esse comando cria um workspace (go.work) que inclui os diretórios ./math e ./sistema.

✅ Workspaces permitem que você gerencie múltiplos módulos Go ao mesmo tempo, facilitando o desenvolvimento de projetos que dependem de múltiplos módulos locais.


6️⃣ go mod tidy -e

🔹 O que faz?

O comando go mod tidy remove dependências não utilizadas e adiciona as necessárias ao go.mod e go.sum. O flag -e faz com que o Go ignore erros, permitindo que o comando continue mesmo que existam problemas com dependências.

✅ Isso ajuda a manter o go.mod limpo e atualizado sem interromper o fluxo de trabalho caso existam problemas menores.

[⬅ Voltar para o README principal](/README.MD)



