package main

import (
	"net/http"

	"github.com/fullcycle/curso-go/7-Apis/configs"
	"github.com/fullcycle/curso-go/7-Apis/internal/entity"
	"github.com/fullcycle/curso-go/7-Apis/internal/infra/database"
	"github.com/fullcycle/curso-go/7-Apis/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // Driver SQLite sem CGO // Importa o driver sem CGO
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// db, err := gorm.Open(sqlite.Dialector{DSN: "test.db"}, &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
