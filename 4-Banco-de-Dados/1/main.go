package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

/***
1) - docker-compose up -d
2) - docker-compose exec mysql bash || docker exec -it mysql bash || docker-compose down && docker-compose up -d || docker logs mysql
3) - mysql -u root -p goexpert
4) - create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));
***/

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)") // $1 $2 $3 caso utilize sqlite
	if err != nil {
		return err
	}
	defer stmt.Close()
	// result ele contempla os dados que foram inseridos
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	return err
}
