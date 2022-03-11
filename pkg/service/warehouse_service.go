package service

import (
	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/repository"
)

type WarehouseService struct {
	repo repository.Warehouse
}

func NewWarehouseService(repo repository.Warehouse) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) Add_warehouse_record(warehouse_data agregator.Warehouse) (int, error) {
	return s.repo.Add_warehouse_record(warehouse_data)
}

func (s *WarehouseService) Show_warehouse_records() ([]agregator.Warehouse, error) {
	return s.repo.Show_warehouse_records()
}
