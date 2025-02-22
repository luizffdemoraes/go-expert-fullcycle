**Comandos Docker e Mysql.:**

1️⃣ docker-compose down && docker-compose up -d

🔹 O que faz?

🔹 docker-compose down

Para e remove todos os containers, redes e volumes criados pelo docker-compose up.
Ele encerra completamente os serviços definidos no docker-compose.yml.

🔹 docker-compose up -d

Cria e inicia os containers novamente com base no docker-compose.yml.
A flag -d executa os containers em segundo plano (modo detached).


2️⃣ docker-compose exec mysql bash || docker exec -it mysql bash 

🔹 O que faz?

🔹 docker exec -it mysql bash

Executa um shell (bash) dentro do container mysql definido no docker-compose.yml.
Requer que o docker-compose esteja sendo usado para gerenciar o container.

🔹 docker exec -it mysql bash 

Executa um shell (bash) dentro do container mysql usando o docker diretamente (sem docker-compose).
Útil se o container foi iniciado manualmente via docker run em vez de docker-compose.


3️⃣ mysql -u root -p goexpert

🔹 O que faz?

🔹 mysql

Inicia o cliente MySQL no terminal.

🔹 -u root

Especifica o usuário root para conexão.

🔹 -p

Pede a senha do usuário antes de conectar.

🔹 goexpert

Nome do banco de dados ao qual deseja se conectar.


4️⃣ create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));

🔹 O que faz?

🔹 CREATE TABLE products

Cria uma nova tabela chamada products no banco de dados.

🔹 id VARCHAR(255)

Cria a coluna id do tipo texto (máximo de 255 caracteres).

🔹 name VARCHAR(80)

Cria a coluna name para armazenar nomes de produtos com até 80 caracteres.

🔹 price DECIMAL(10,2)

Cria a coluna price para armazenar valores numéricos com 10 dígitos no total e 2 casas decimais (exemplo: 99999999.99).

🔹 PRIMARY KEY (id)

Define a coluna id como chave primária, garantindo que cada produto tenha um identificador único.

🔹 O que acontece ao executar?

O banco de dados cria a tabela products com as colunas definidas.
O campo id será único para cada produto, impedindo valores duplicados.

5️⃣ SHOW TABLES products;

🔹 O que faz?

Isso retorna uma lista de todas as tabelas existentes no banco.

6️⃣ DROP TABLE products;

🔹 O que faz?

Remove completamente a tabela products do banco de dados.

7️⃣ TRUNCATE TABLE products;

🔹 O que faz?

Remove todos os registros da tabela products, mas mantém a estrutura da tabela.


[⬅ Voltar para o README principal](/README.MD)



