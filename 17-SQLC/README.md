# 17-SQLC

Referências e ferramentas utilizadas neste módulo:

- **[sqlc](https://sqlc.dev/)** — Compile SQL to type-safe code
- **[golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)** — CLI de migrações para Go

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
