package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lffm1994/17-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	db      *sql.DB
	queries *db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		db:      dbConn,
		queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("error on rollback: %w", err)
		}
		return err
	}

	return tx.Commit()
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)
}
