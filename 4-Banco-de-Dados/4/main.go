package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

/***
https://gorm.io/
DROP TABLE products;

Em GORM, a associação Has One indica que um registro em um modelo possui exatamente um registro associado em outro modelo.
Isso é usado quando, por exemplo, um usuário tem um único perfil ou uma entidade tem um detalhe exclusivo relacionado.
***/

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Criação da tabela
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Criação de categoria
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Criação de produto
	db.Create(&Product{
		Name:       "Mouse",
		Price:      80.0,
		CategoryID: 1,
	})

	// Criação de serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
