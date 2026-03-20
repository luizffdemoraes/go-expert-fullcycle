package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       string // decimal(10,2) no MySQL — use string no driver
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
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

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			CategoryID:  argsCategory.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	// queries := db.New(dbConn)

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go Course", Valid: true},
		Price:       "99.90",
	}
	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Course", Valid: true},
	}

	courseDB := NewCourseDB(dbConn)
	err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err)
	}
}
