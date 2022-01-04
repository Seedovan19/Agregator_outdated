package repository

import (
	agregator "github.com/seedovan19/Agregator"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user agregator.User) (int, error)
}

type Warehouse interface {
}

type Repository struct {
	Authorization
	Warehouse
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
