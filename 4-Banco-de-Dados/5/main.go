package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
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
	db.AutoMigrate(&Product{}, &Category{})

	// Criação de categoria
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	// Criação de produto
	db.Create(&Product{
		Name:       "Panela",
		Price:      180.0,
		CategoryID: 2,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(product.Name, category.Name)
		}
	}
}
