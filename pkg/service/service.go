package service

import (
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/repository"
)

type Authorization interface {
	CreateUser(user agregator.User) (int, error)
}

type Warehouse interface {
}

type Service struct {
	Authorization
	Warehouse
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
