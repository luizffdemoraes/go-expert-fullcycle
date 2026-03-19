package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lffm1994/17-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)

	/*** Create Category
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:   uuid.New().String(),
		Name: "Category 1",
		Description: sql.NullString{
			String: "Description 1",
			Valid:  true,
		},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description)
	}
	***/

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:   "78e2d7a2-5f73-4bdc-b37b-205ff12c5e0b",
		Name: "Backend updated",
		Description: sql.NullString{
			String: "Backend updated description",
			Valid:  true,
		},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description)
	}

	err = queries.DeleteCategory(ctx, "78e2d7a2-5f73-4bdc-b37b-205ff12c5e0b")
	if err != nil {
		panic(err)
	}
}
