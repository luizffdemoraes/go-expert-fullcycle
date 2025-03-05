package database

import "github.com/fullcycle/curso-go/7-Apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
