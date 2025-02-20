package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey;autoIncrement"`
	Name string
}

type Product struct {
	gorm.Model
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	CategoryID int
	Category   Category
	Price      float64
}

/***
https://gorm.io/
DROP TABLE products;

Em GORM (ORM do Go), Belongs To é uma associação que indica que um modelo pertence a outro.
Isso é usado quando um registro em uma tabela faz referência a um único registro em outra tabela.
***/

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Criação da tabela
	// db.AutoMigrate(&Product{}, &Category{})

	// Criação de categoria
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	db.Create(&Product{
		Name:       "Mouse",
		Price:      80.0,
		CategoryID: 1,
	})

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
