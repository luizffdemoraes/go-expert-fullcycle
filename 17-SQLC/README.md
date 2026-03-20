# 17-SQLC

Referências e ferramentas utilizadas neste módulo:

- **[sqlc](https://sqlc.dev/)** — Compile SQL to type-safe code
- **[golang-migrate/migrate](https://github.com/golang-migrate/migrate)** — CLI de migrações para Go
- **[jmoiron/sqlx](https://github.com/jmoiron/sqlx)** — Extensões para `database/sql` (Go)

---

## Processo de instalação do migrate (passo a passo)

### Passo 1: Remover configurações anteriores
```bash
sudo rm -f /etc/apt/sources.list.d/migrate.list
```

### Passo 2: Definir a versão
```bash
VERSION="v4.18.2"
```

### Passo 3: Baixar e extrair
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/$VERSION/migrate.linux-amd64.tar.gz | tar xvz
```

### Passo 4: Mover para o PATH
```bash
sudo mv migrate /usr/local/bin/
```

### Passo 5: Verificar instalação
```bash
migrate -version
```

---

## Criação e limpeza das migrações

### Criar primeira migração
```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```

### Executar migrações
```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
```

### Rodar migrações via Makefile
> O `make migrate-up` e `make migrate-down` garantem que o banco `courses` exista (criando via Docker) antes de executar o `migrate`.

```bash
make create-migration
make migrate-up
make migrate-down
```

Observacao: ao rodar `make migrate-down`, o `migrate` pede confirmação:
```
Are you sure you want to apply all down migrations? [y/N]
```
Digite `y` e pressione Enter para continuar.

### Acessar o container do MySQL e o banco
```bash
docker compose exec mysql bash
mysql -uroot -proot courses
```

Arquivos gerados (exemplo):
```text
/home/lffm1994/sql/migrations/000001_init.up.sql
/home/lffm1994/sql/migrations/000001_init.down.sql
```

### Remover migrações (limpar diretório)
```bash
rm -rf ~/sql/migrations
rm -rf ~/sql
```

---

## Vendo a transação funcionar

1) Rode o exemplo de transação:
```bash
go run ./cmd/server-transaction/main.go
```

2) Valide no MySQL que as tabelas foram populadas:
```bash
docker compose exec mysql mysql -uroot -proot -D courses -e "
SHOW TABLES;
SELECT COUNT(*) AS categories_count FROM categories;
SELECT COUNT(*) AS courses_count FROM courses;
SELECT c.name AS category, co.name AS course, co.price
FROM categories c
JOIN courses co ON co.category_id = c.id
ORDER BY co.id DESC
LIMIT 5;
"
```

Se a transação estiver funcionando, o resultado deve mostrar:
- `categories_count > 0`
- `courses_count > 0`
- `price` preenchido (ex: `99.90`).
