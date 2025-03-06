package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/fullcycle/curso-go/7-Apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // Driver SQLite sem CGO
)

func TestCreateProduct(t *testing.T) {

	db, err := gorm.Open(sqlite.Dialector{DriverName: "sqlite", DSN: "file::memory:"}, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)

	err = productDB.Create(product)

	assert.NoError(t, err)
	assert.Equal(t, product.ID, product.ID)
	assert.Equal(t, product.Name, product.Name)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Dialector{DriverName: "sqlite", DSN: "file::memory:"}, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Dialector{DriverName: "sqlite", DSN: "file::memory:"}, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	// Realiza a migração
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	// Cria um novo produto
	product, err := entity.NewProduct("Produto 1", 120.0)
	assert.NoError(t, err)

	// Salva no banco
	if err := db.Create(product).Error; err != nil {
		t.Error(err)
	}

	// Instancia o repositório e busca o produto
	productDB := NewProduct(db)
	productFound, err := productDB.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
}

func TestUpdateProductByID(t *testing.T) {

	db, err := gorm.Open(sqlite.Dialector{DriverName: "sqlite", DSN: "file::memory:"}, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", productFound.Name)
}

func TestDeleteProductByID(t *testing.T) {

	db, err := gorm.Open(sqlite.Dialector{DriverName: "sqlite", DSN: "file::memory:"}, &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product)
	assert.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Empty(t, product)
}
