package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Realizando conexão com o banco
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
	product.Price = 100.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selesctProduct(ctx, db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
}

func insertProduct(db *sql.DB, product *Product) error {
	// Preparando a query
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)") // $1 $2 $3 caso utilize sqlite
	if err != nil {
		return err
	}
	defer stmt.Close()
	// result ele contempla os dados que foram inseridos
	// Executando a query
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	// Preparando a query
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?") // $1 $2 $3 caso utilize sqlite
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Executando a query
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	return err
}

func selesctProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	// Preparando a query
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?") // $1 $2 $3 caso utilize sqlite
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	// Executando a query
	stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	// err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
