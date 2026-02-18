# 13-GRAPHQL

Projeto de API GraphQL em Go utilizando [gqlgen](https://gqlgen.com/) e SQLite.

## Pré-requisitos

- Go 1.25+
- SQLite3 (para criação do banco de dados)

## Como criar o projeto do zero

Siga os passos abaixo para inicializar um novo projeto GraphQL com gqlgen:

```bash
# 1. Inicializar o módulo Go
go mod init github.com/seu-usuario/13-GraphQL

# 2. Inicializar o gqlgen (cria a estrutura base do schema e resolvers)
go run github.com/99designs/gqlgen init

# 3. Gerar o código a partir do schema GraphQL (após editar graph/schema.graphqls)
go run github.com/99designs/gqlgen generate
```

> **Nota:** Ajuste o caminho do módulo em `go mod init` conforme seu repositório. O comando `gqlgen init` cria os arquivos `graph/schema.graphqls`, `graph/schema.resolvers.go`, `gqlgen.yml`, entre outros. Use `gqlgen generate` sempre que alterar o schema.

## Configuração do banco de dados

O projeto usa SQLite com o arquivo `data.db`. Crie as tabelas necessárias:

```bash
sqlite3 data.db
```

No prompt do SQLite:

```sql
CREATE TABLE categories (id TEXT, name TEXT, description TEXT);
CREATE TABLE courses (id string, name string, description string, category_id string);
```

Ou em uma linha (no terminal):

```bash
sqlite3 data.db "CREATE TABLE categories (id TEXT, name TEXT, description TEXT);"
sqlite3 data.db "CREATE TABLE courses (id string, name string, description string, category_id string);"
```

> Em SQLite, o tipo recomendado para strings é `TEXT`. O arquivo `data.db` será criado no diretório de onde você executar o servidor (por padrão na pasta do projeto).

## Executando o servidor

```bash
go run ./cmd/server
```

O servidor sobe na porta **8080**. Acesse o **GraphQL Playground** em:

- **http://localhost:8080/**

As requisições GraphQL são enviadas para o endpoint **http://localhost:8080/query**.

## Uso da API GraphQL

### Criar uma categoria (Mutation)

No Playground (ou em qualquer cliente GraphQL), execute:

```graphql
mutation createCategory {
  createCategory(input: {
    name: "Tecnologia",
    description: "Curso de Tecnologia",
  }) {
    id
    name
    description
  }
}
```

Isso insere uma nova categoria no banco e retorna `id`, `name` e `description` da categoria criada.

### Listar categorias (Query)

Para buscar todas as categorias:

```graphql
query queryCategories {
  categories {
    id
    name
    description
  }
}
```

A resposta será uma lista com todas as categorias cadastradas.

### Criar um curso (Mutation)

Para criar um curso vinculado a uma categoria, use o `id` de uma categoria existente (por exemplo, retornado por `createCategory` ou `queryCategories`):

```graphql
mutation createCourse {
  createCourse(input: {
    name: "Full Cycle",
    description: "The best!",
    categoryId: "8bcd0e47-bda1-4d7f-a0a5-fe1472ef676b"
  }) {
    id
    name
  }
}
```

Substitua `categoryId` pelo ID real de uma categoria cadastrada no banco.

### Listar cursos (Query)

Para buscar todos os cursos:

```graphql
query queryCourses {
  courses {
    id
    name
  }
}
```

A resposta será uma lista com todos os cursos cadastrados.

### Categorias com cursos (Query aninhada)

A API permite buscar categorias já trazendo os cursos de cada uma em uma única query:

```graphql
query queryCategoriesWithCourses {
  categories {
    id
    name
    courses {
      id
      name
      description
    }
  }
}
```

Cada categoria retornada inclui o campo `courses` com a lista de cursos vinculados a ela.

#### O que foi necessário para funcionar

1. **`gqlgen.yml`** – Configuração dos models para usar structs customizados em `graph/model` (Category e Course). O struct `Category` em Go não possui o campo `courses`; ao mapear o tipo GraphQL `Category` para esse model, o gqlgen passa a gerar a interface `CategoryResolver` com o método `Courses(ctx, obj)`, que você implementa para carregar os cursos por `categoryId` no banco.

   Trecho relevante em `gqlgen.yml`:

   ```yaml
   models:
     Category:
       model:
         - github.com/luizffdemoraes/13-GraphQL/graph/model.Category
     Course:
       model:
         - github.com/luizffdemoraes/13-GraphQL/graph/model.Course
   ```

2. **Implementação do resolver** – No `graph/schema.resolvers.go`, o resolver `Category.Courses` chama o banco (ex.: `CourseDB.FindByCategoryID(obj.ID)`) e retorna a lista de cursos daquela categoria.

3. **Regenerar o código** – Após alterar o schema ou o `gqlgen.yml`, é preciso rodar novamente:

   ```bash
   go run github.com/99designs/gqlgen generate
   ```

   Assim o gqlgen regera `generated.go` e os stubs dos resolvers (como `Category.Courses`) para a implementação atual.

## Fluxo resumido

1. **Criar o projeto:** `go mod init` → `go run github.com/99designs/gqlgen init` → editar o schema → `go run github.com/99designs/gqlgen generate`
2. **Preparar o banco:** criar o arquivo `data.db` e as tabelas `categories` e `courses` com `sqlite3`
3. **Subir o servidor:** `go run ./cmd/server`
4. **Testar:** abrir http://localhost:8080/ e usar as mutations e queries acima

## Estrutura principal

- `graph/schema.graphqls` – definição do schema GraphQL (tipos, queries, mutations)
- `graph/schema.resolvers.go` – implementação dos resolvers (lógica das queries e mutations)
- `internal/database/` – acesso ao SQLite (ex.: categorias)
- `cmd/server/` – ponto de entrada do servidor HTTP e GraphQL

## Links úteis

- **[gqlgen](https://gqlgen.com/)** – biblioteca oficial para construir servidores GraphQL em Go com abordagem schema-first, type-safety e codegen. Documentação, getting started e referência da API.
