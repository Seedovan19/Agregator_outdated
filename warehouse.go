package agregator

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model

	ID                 int64  `gorm:"primaryKey" json:"id"`
	Name               string `gorm:"not null" json:"name"`
	Square             int64  `gorm:"not null" json:"square"`
	Adress             string `gorm:"not null" json:"adress"`
	Shelf_storage_cost int64  `gorm:"not null" json:"shelf_storage"`
	Floor_storage_cost int64  `gorm:"not null" json:"floor_storage"`
	Description        string `json:"description"`
	Comment            string `json:"comment"`
	// Building           Building `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type Building struct {
// 	gorm.Model

// 	Warehouse_class      string    `gorm:"not null" json:"class"`
// 	Year_of_construction time.Time `json:"year_of_construction"`
// }
