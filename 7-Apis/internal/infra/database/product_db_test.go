package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/fullcycle/curso-go/7-Apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // Importa o driver sem CGO
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
