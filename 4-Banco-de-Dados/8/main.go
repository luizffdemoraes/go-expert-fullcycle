package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


// Definição da struct Product (produto) para representar a tabela no banco
type Product struct {
	ID       uint
	Name     string
	Price    float64
	CategoryID uint
}

// Definição da struct Category (categoria) para representar a tabela no banco
type Category struct {
	ID   uint
	Name string
}

func main() {
	// Definição da string de conexão (DSN - Data Source Name)
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	// Conexão com o banco de dados usando GORM e o driver MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // Encerra o programa se houver erro na conexão
	}

	// Criação automática das tabelas 'products' e 'categories' se não existirem
	db.AutoMigrate(&Product{}, &Category{})

	// Inicia uma transação no banco de dados
	tx := db.Begin()

	var c Category
	// Realiza um SELECT na tabela 'categories' buscando o registro com ID 1,
	// usando um bloqueio de linha "FOR UPDATE" para evitar concorrência
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err) // Encerra o programa se a consulta falhar
	}

	// Altera o nome da categoria para "Eletronicos"
	c.Name = "Eletronicos"

	// Salva a alteração no banco dentro da transação
	tx.Debug().Save(&c)

	// Confirma a transação (commit), aplicando as mudanças permanentemente
	tx.Commit()
}
