package service

//go:generate mockgen -source=service.go -destination=mocks/mock.go

import (
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/repository"
)

type Warehouse interface {
	Add_warehouse_record(warehouse_data agregator.Warehouse, building_data agregator.Building) (int, error)
	Show_warehouse_records() ([]agregator.Warehouse, error)
}

type Service struct {
	Warehouse
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Warehouse: NewWarehouseService(repos.Warehouse),
	}
}
