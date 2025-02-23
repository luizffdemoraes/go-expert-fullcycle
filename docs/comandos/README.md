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


7️⃣ go test .

🔹 O que faz?

Executa os testes dentro do pacote atual (. significa "pacote atual").

✅ O Go procura por arquivos com _test.go, executa as funções Test*, e exibe o resultado.


8️⃣ go test -v

🔹 O que faz?

Executa os testes de forma verbosa, mostrando detalhes sobre cada teste rodado.

✅ Exibe mensagens como:

9️⃣ go test -coverprofile=coverage

🔹 O que faz?

Executa os testes e gera um relatório de cobertura de código no arquivo coverage.

✅ Esse arquivo pode ser usado para analisar quais partes do código foram testadas.

🔟 go tool cover -html=coverage

🔹 O que faz?

Abre um relatório visual em HTML mostrando a cobertura do código.

✅ Exibe quais linhas do código foram testadas e quais não foram, destacando o código em cores diferentes.

1️⃣1️⃣ go test -bench=. -count=10 -benchtime=3s -benchmem

🔹 O que faz?

Executa testes de benchmark para medir o desempenho do código no pacote atual.

✅ Diferente dos testes normais (go test), que verificam a corretude do código, o benchmark mede o tempo de execução de funções específicas.

1️⃣2️⃣ go test -bench=. -run=^#

🔹 O que faz?

Executa apenas os benchmarks, ignorando os testes unitários.

✅ Explicação:

-bench=. → Roda todos os benchmarks no pacote atual.
-run=^# → Regex que não corresponde a nenhum teste unitário, ou seja, impede que os testes normais sejam executados.

1️⃣3️⃣ go test -bench=. -count=10 -benchtime=3s

🔹 O que faz?

Executa benchmarks 10 vezes, rodando cada um por pelo menos 3 segundos.

✅ Explicação dos flags:

-bench=. → Roda todos os benchmarks no pacote atual.
-count=10 → Executa cada benchmark 10 vezes, garantindo resultados mais estáveis.
-benchtime=3s → Faz cada benchmark rodar por pelo menos 3 segundos, ajustando dinamicamente o número de iterações.

1️⃣4️⃣ go test -bench=. -benchmem

🔹 O que faz?

🔹 Executa benchmarks e mede o uso de memória.

✅ Explicação dos flags:

-bench=. → Roda todos os benchmarks.
-benchmem → Mede o uso de memória, incluindo:
Bytes alocados por operação (B/op)
Número de alocações de heap (allocs/op)



[⬅ Voltar para o README principal](/README.MD)



