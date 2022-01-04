package repository

import (
	agregator "github.com/seedovan19/Agregator"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user agregator.User) (int, error) {
	r.db.Create(&user)
	return 0, nil
}
