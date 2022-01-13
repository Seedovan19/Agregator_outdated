package service

//go:generate mockgen -source=service.go -destination=mocks/mock.go

import (
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/repository"
)

type Authorization interface {
	CreateUser(user agregator.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Warehouse interface {
	Add_warehouse_record(warehouse_data agregator.Warehouse, building_data agregator.Building) (int, error)
	Show_warehouse_records() ([]agregator.Warehouse, error)
}

type Service struct {
	Authorization
	Warehouse
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Warehouse:     NewWarehouseService(repos.Warehouse),
	}
}
