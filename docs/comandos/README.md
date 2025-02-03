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

[⬅ Voltar para o README principal](../../README.md)



