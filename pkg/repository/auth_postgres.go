package repository

import (
	"log"

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

func (r *AuthPostgres) GetUser(email, password string) (agregator.User, error) {
	var user agregator.User

	if result := r.db.Where("email = ? and password = ?", email, password).First(&user); result.Error != nil {
		return user, result.Error
	} else {
		rows, err := result.Rows()
		if err != nil {
			log.Fatalf(err.Error())
		}
		rows.Scan(&user.ID, &user.Email, &user.Password)

		return user, nil
	}
}
