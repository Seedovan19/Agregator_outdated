package agregator

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model

	ID                 int64  `gorm:"primaryKey"`
	Name               string `gorm:"not null"`
	Square             int64  `gorm:"not null" `
	Adress             string `gorm:"not null"`
	Shelf_storage_cost int64  `gorm:"not null"`
	Floor_storage_cost int64  `gorm:"not null"`
	Description        string
	Comment            string
	Building           Building `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Phone              string
}

type Building struct {
	gorm.Model

	Warehouse_class      string `gorm:"not null"`
	Year_of_construction time.Time
}
