package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Price float64
}

// https://gorm.io/
// DROP TABLE products;
func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Criação da tabela
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name: "Notebook",
	// 	Price: 1000.00,
	// })

	// create batch
	// 	products := []Product{
	// 		{Name: "Netbook", Price: 1000.00},
	// 		{Name: "Mouse", Price: 50.00},
	// 		{Name: "Keyboard", Price: 100.00},
	// 	}
	// 	db.Create(&products)

	//select one
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

}
