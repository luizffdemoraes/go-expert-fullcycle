# 15-UPLOAD-S3 — Upload de arquivos para AWS S3

Projeto em Go para envio de arquivos a um bucket S3 (ou serviço compatível como MinIO), evoluindo de uma implementação serial até uploads concorrentes com limite de goroutines e retentativas em caso de erro.

---

## Índice

1. [Entendendo o problema referente a upload](#1-entendendo-o-problema-referente-a-upload)
2. [Gerando arquivos exemplo](#2-gerando-arquivos-exemplo)
3. [Configurando AWS session](#3-configurando-aws-session)
4. [Desenvolvendo função de upload](#4-desenvolvendo-função-de-upload)
5. [Finalizando primeira implementação](#5-finalizando-primeira-implementação)
6. [Criando credenciais na AWS](#6-criando-credenciais-na-aws)
7. [Fazendo upload de forma serial](#7-fazendo-upload-de-forma-serial)
8. [Realizando uploads usando goroutines](#8-realizando-uploads-usando-goroutines)
9. [Limitando quantidade máxima de upload](#9-limitando-quantidade-máxima-de-upload)
10. [Fazendo retentativas de erro](#10-fazendo-retentativas-de-erro)

---

## 1. Entendendo o problema referente a upload

**Objetivo:** Enviar vários arquivos de um diretório local para um bucket S3.

- **Entrada:** pasta (ex.: `tmp/`) com arquivos gerados ou copiados.
- **Saída:** os mesmos arquivos disponíveis no S3, identificados pela mesma chave (nome do arquivo).
- **Requisitos que evoluímos ao longo do projeto:**
  - Funcionar com AWS S3 ou com MinIO (S3-compatível).
  - Não bloquear em um único arquivo (concorrência).
  - Não sobrecarregar rede/S3 (limite de uploads simultâneos).
  - Reenviar arquivos que falharem (abertura ou PutObject).

**Onde isso aparece no código:** toda a lógica está em `cmd/uploader/main.go`: leitura do diretório, configuração do cliente S3 e chamadas a `uploadFile`.

---

## 2. Gerando arquivos exemplo

Para testar o upload sem depender de arquivos reais, o projeto inclui um **generator** que cria milhares de arquivos de texto em `./tmp/`.

**Implementação:** `cmd/generator/main.go`

- Loop de `0` a `4999` (5000 arquivos).
- Cada arquivo: `./tmp/file_0.txt`, `./tmp/file_1.txt`, … com conteúdo `"Hello, World!"`.
- Uso de `os.Create`, `WriteString` e `Close` para criar os arquivos.

**Como rodar:**

```bash
cd 15-UPLOAD-S3
go run ./cmd/generator
```

Gera os arquivos em `./tmp/`. Rode o generator **antes** do uploader se a pasta `tmp` estiver vazia ou não existir.

---

## 3. Configurando AWS session

A “sessão” com a AWS (ou MinIO) é a configuração do cliente S3: região, credenciais e endpoint. Isso é feito no `init()` do uploader para que esteja pronto antes do `main()`.

**Implementação:** `cmd/uploader/main.go` — bloco `init()`.

- **Variáveis de ambiente:** carregadas com `godotenv` a partir de um `.env` (procurando na pasta atual e nas pastas pai, para funcionar com qualquer cwd).
- **Região:** `AWS_REGION` (default `us-east-1`).
- **Credenciais:** `AWS_ACCESS_KEY_ID` e `AWS_SECRET_ACCESS_KEY` via `credentials.NewStaticCredentialsProvider`.
- **Endpoint customizado:** `S3_ENDPOINT` (default `http://localhost:9000`) e `UsePathStyle: true` para MinIO e outros S3-compatíveis.
- **Bucket e pasta:** `S3_BUCKET` (default `goexpert-bucket-exemplo`) e `UPLOAD_DIR` (default `tmp`), com caminho resolvido a partir da raiz do projeto (onde está o `go.mod`).

Assim, a “sessão” é o cliente S3 configurado com `config.LoadDefaultConfig` + `s3.NewFromConfig` e as opções de endpoint.

---

## 4. Desenvolvendo função de upload

A função que efetivamente envia um arquivo para o S3 é `uploadFile`.

**Implementação:** `cmd/uploader/main.go` — função `uploadFile`.

- **Parâmetros:** nome do arquivo, canal de controle de slots (`uploadControl`) e canal de retry (`errFileUpload`).
- **Fluxo:**
  1. Monta o caminho completo com `filepath.Join(uploadDir, fileName)`.
  2. Abre o arquivo com `os.Open`.
  3. Se falhar ao abrir: libera o slot, envia o `fileName` para `errFileUpload` (retry) e retorna.
  4. Chama `s3Client.PutObject` com `Bucket`, `Key` (nome do arquivo) e `Body` (o arquivo).
  5. Se `PutObject` falhar: libera o slot, envia o `fileName` para retry e retorna.
  6. Em sucesso: libera o slot e retorna.

O “upload” em si é o `PutObject`; a abertura do arquivo e o tratamento de erro fazem parte da mesma função para centralizar a lógica.

---

## 5. Finalizando primeira implementação

A “primeira implementação” é o fluxo mínimo: configurar o cliente, ler o diretório e enviar cada arquivo. No código atual isso está integrado com goroutines, limite de concorrência e retry; a ideia de “finalizar” é ter:

- `init()` configurando o cliente S3 e variáveis globais (`uploadDir`, `S3Bucket`).
- `main()` abrindo o diretório de upload, iterando com `Readdir(1)` e chamando a função de upload para cada arquivo.
- `uploadFile` abrindo o arquivo e chamando `PutObject`.

Com isso o programa já sobe todos os arquivos de `tmp/` para o bucket. Os próximos tópicos evoluem para serial explícito, concorrência, limite e retry.

---

## 6. Criando credenciais na AWS

Para AWS real (não MinIO), é preciso usar credenciais IAM (Access Key + Secret).

- **Console AWS:** IAM → Users → Security credentials → Create access key (ex.: “Application running outside AWS”).
- Guarde **Access Key ID** e **Secret Access Key** e coloque no `.env`:

```env
AWS_ACCESS_KEY_ID=sua_access_key
AWS_SECRET_ACCESS_KEY=sua_secret_key
AWS_REGION=us-east-1
# Deixe S3_ENDPOINT em branco ou comente para usar o endpoint padrão da AWS
S3_BUCKET=seu-bucket
```

Para **MinIO** (desenvolvimento local), use o `.env.example` do projeto:

```bash
cp .env.example .env
# Ajuste se necessário; padrão: minioadmin/minioadmin, endpoint http://localhost:9000
```

**Onde é usado no código:** `init()` em `cmd/uploader/main.go` lê essas variáveis via `env("AWS_ACCESS_KEY_ID", "minioadmin")` etc. e passa para `credentials.NewStaticCredentialsProvider`. Não commite o `.env` com credenciais reais.

---

## 7. Fazendo upload de forma serial

Upload **serial** significa enviar um arquivo por vez, sem goroutines: ler um arquivo do diretório, chamar a função de upload (síncrono), depois o próximo.

**Como seria no código (conceitual):**

- Em `main()`, em vez de `go uploadFile(...)`, chamar `uploadFile(...)` diretamente (e sem canais de controle/retry).
- A função de upload não precisaria de `uploadControl` nem `errFileUpload`; apenas abriria o arquivo e chamaria `PutObject`.

No código atual, o fluxo já está paralelo; a “versão serial” é o passo anterior na evolução: um `for` no `main` com uma chamada síncrona por arquivo. É mais simples, mas mais lenta que a versão com goroutines.

---

## 8. Realizando uploads usando goroutines

Para ganhar desempenho, cada arquivo é enviado em uma **goroutine** separada, em paralelo.

**Implementação:** `cmd/uploader/main.go` — `main()`.

- `wg sync.WaitGroup` (global) para esperar todas as goroutines.
- Para cada arquivo retornado por `Readdir(1)`:
  - `wg.Add(1)`.
  - `go uploadFile(files[0].Name(), uploadControl, errFileUpload)`.
- No fim do loop: `wg.Wait()` para não encerrar o programa antes de todos os uploads terminarem.
- Em `uploadFile`, `defer wg.Done()` garante que cada goroutine decremente o WaitGroup ao terminar (sucesso ou erro).

Assim, vários arquivos são enviados ao mesmo tempo, cada um em uma goroutine.

---

## 9. Limitando quantidade máxima de upload

Se dispararmos uma goroutine por arquivo sem limite, podemos abrir muitas conexões e sobrecarregar rede ou S3. Por isso usamos um **semáforo** com canal: no máximo N uploads simultâneos.

**Implementação:** `cmd/uploader/main.go`.

- **Canal de controle:** `uploadControl := make(chan struct{}, 100)` — buffer de 100 “slots”.
- **Antes de iniciar um upload:** `uploadControl <- struct{}{}` (bloqueia se já houver 100 uploads em andamento).
- **Depois de terminar um upload (sucesso ou erro):** `<-uploadControl` (libera um slot).

Assim, nunca há mais do que 100 uploads ativos. O tamanho do buffer (100) é o “limite máximo de uploads simultâneos”. A struct vazia é usada só como unidade de sincronização (baixo custo de memória).

---

## 10. Fazendo retentativas de erro

Quando um upload falha (erro ao abrir o arquivo ou erro no `PutObject`), o arquivo é recolocado em uma fila para nova tentativa, em vez de ser descartado.

**Implementação:** `cmd/uploader/main.go`.

- **Canal de retry:** `errFileUpload := make(chan string, 100)` — recebe o **nome do arquivo** a ser reenviado.
- **Em `uploadFile`:** em caso de erro (abertura ou PutObject), após liberar o slot fazemos `errFileUpload <- fileName`.
- **Goroutine de retry:** uma goroutine fica em loop lendo de `errFileUpload` (com `select` e receive em dois valores para detectar canal fechado). Para cada nome recebido:
  - Obtém um slot: `uploadControl <- struct{}{}`.
  - `wg.Add(1)` e `go uploadFile(fileName, uploadControl, errFileUpload)`.

Assim, qualquer falha (abertura ou S3) recoloca o mesmo arquivo na fila para uma nova tentativa. O encerramento correto é: após o loop principal, `wg.Wait()` e em seguida `close(errFileUpload)` para a goroutine de retry sair do loop (quando `ok == false` no receive).

---

## Estrutura do projeto

```
15-UPLOAD-S3/
├── cmd/
│   ├── generator/
│   │   └── main.go    # Gera arquivos em ./tmp/
│   └── uploader/
│       └── main.go    # Upload para S3 com concorrência, limite e retry
├── tmp/               # Arquivos gerados (e lidos pelo uploader)
├── .env.example       # Exemplo de variáveis para MinIO/AWS
├── go.mod
├── go.sum
└── README.md
```

---

## Como executar

1. **Gerar arquivos (opcional):**
   ```bash
   go run ./cmd/generator
   ```

2. **Configurar ambiente:**
   ```bash
   cp .env.example .env
   # Edite .env com credenciais e bucket (ou use padrão MinIO)
   ```

3. **Subir o uploader:**
   ```bash
   go run ./cmd/uploader
   ```

Para AWS real, configure no `.env` as credenciais e o bucket; para MinIO local, use o endpoint (ex.: `http://localhost:9000`) e crie o bucket no MinIO antes de rodar o uploader.

---

## Dependências principais

- `github.com/aws/aws-sdk-go-v2` e `service/s3` — cliente S3.
- `github.com/aws/aws-sdk-go-v2/config` e `credentials` — configuração e credenciais.
- `github.com/joho/godotenv` — carregamento do `.env`.

Todas listadas em `go.mod`; use `go mod tidy` se precisar atualizar.
