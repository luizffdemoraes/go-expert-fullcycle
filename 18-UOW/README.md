# 18-UOW — Unit of Work

## O que é o Unit of Work (UoW)

**Unit of Work** é um padrão de projeto que agrupa várias operações de leitura e escrita no banco em **uma única unidade transacional**. Em vez de cada repositório abrir e confirmar sua própria transação, o UoW coordena **um** `beginTransaction` / `commit` / `rollback` para todo o conjunto de alterações.

## Para que serve

- **Orquestrar uma transação** envolvendo vários repositórios (por exemplo: inserir categoria e curso na mesma transação).
- **Garantir que todos os repositórios usem o mesmo contexto** (no caso deste projeto, o mesmo `*sql.Tx`), para que o banco enxergue as mudanças como um bloco único.
- **Centralizar o registro** de fábricas de repositórios ligadas à transação atual (`Register` / `GetRepository`).

## O que o UoW corrige (e o que não corrige)

**Corrige / endereça:**

- **Atomicidade**: ou todas as operações dentro do `Do` são confirmadas, ou nenhuma (rollback em caso de erro).
- **Estado inconsistente parcial**: evita situações em que, por exemplo, a categoria foi gravada e o curso falhou depois, deixando dados órfãos quando isso deveria ser um único fluxo.
- **Uso consistente da transação** entre vários inserts/updates no mesmo fluxo de negócio.

**Não substitui:**

- Regras de negócio incorretas (por exemplo, usar um `category_id` que não existe ou que não corresponde ao registro recém-inserido).
- Modelagem errada do schema ou das queries — isso continua sendo ajustado no domínio e na camada de persistência.

## Quem criou / documentou o padrão

O padrão **Unit of Work** foi **nomeado e documentado** por **Martin Fowler** no livro *Patterns of Enterprise Application Architecture* (Addison-Wesley, 2002). Fowler descreve o UoW como forma de rastrear mudanças em entidades e consolidá-las em um único commit ao banco.

> Referência: Fowler, M. *Patterns of Enterprise Application Architecture*. Capítulo “Unit of Work”.

A implementação em `pkg/uow` deste repositório é código de apoio ao curso/projeto (padrão aplicado em Go com `database/sql`); o **conceito** do padrão vem da literatura de arquitetura de aplicações empresariais acima.
