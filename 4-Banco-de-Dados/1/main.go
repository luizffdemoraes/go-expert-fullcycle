package main

import "github.com/google/uuid"

/***
1) - docker-compose up -d
2) - docker-compose exec mysql bash
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

}
