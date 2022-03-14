package agregator

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model

	ID                 int64        `gorm:"primaryKey" json:"id"`
	Name               string       `gorm:"not null" json:"name"`
	Square             int64        `gorm:"not null" json:"square"`
	Adress             string       `gorm:"not null" json:"adress"`
	Shelf_storage_cost int64        `gorm:"not null" json:"shelf_storage"`
	Floor_storage_cost int64        `gorm:"not null" json:"floor_storage"`
	Description        string       `json:"description"`
	Class              string       `json:"class"`
	Integrated         bool         `gorm:"not null" json:"integrated_flag"`
	Comment            string       `json:"comment"`
	Storagecond        Storagecond  `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Features           Features     `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Security           Security     `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Services           Services     `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Workinghours       Workinghours `gorm:"foreignkey: id; references: ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Storagecond struct {
	gorm.Model

	Floor_storage      int64 `gorm:"not null" json:"floor_storage"`
	Shelf_storage      int64 `gorm:"not null" json:"shelf_storage"`
	Max_storage_weight int64 `json:"max_storage_weight"`
	Max_storage_height int64 `json:"max_storage_height"`
	// 	Warehouse_class      string    `gorm:"not null" json:"class"`
	// 	Year_of_construction time.Time `json:"year_of_construction"`
}

type Features struct {
	gorm.Model

	Condition         string `gorm:"not null" json:"condition"`
	Freezer           bool   `gorm:"not null" json:"freezer"`
	Freezer_size      int64  `json:"freezer_size"`
	Refrigerator      bool   `gorm:"not null" json:"refrigerator"`
	Refrigerator_size int64  `json:"refrigerator_size"`
	Alcohol           bool   `gorm:"not null" json:"alcohol"`
	Pharmacy          bool   `gorm:"not null" json:"pharmacy"`
	Dangerous         bool   `gorm:"not null" json:"dangerous"`
}

type Security struct {
	gorm.Model

	Security_post         bool `json:"security_post"`
	All_day_security      bool `json:"all_day_security"`
	Video_control         bool `json:"video_control"`
	Magnetic_access_locks bool `json:"magnetic_access_locks"`
	Generator             bool `json:"generator"`
	Alarm_system          bool `json:"alarm_system"`
	Fire_system           bool `json:"fire_system"`
}

type Services struct {
	gorm.Model

	Transport_services bool `gorm:"not null" json:"transport_services"`
	Custom             bool `gorm:"not null" json:"custom"`
	Crossdock          bool `gorm:"not null" json:"crossdock"`
	Palletization      bool `gorm:"not null" json:"palletization"`
	Box_pick           bool `gorm:"not null" json:"box_pick"`
}

type Workinghours struct {
	gorm.Model

	Timefrom int64 `gorm:"not null" json:"timefrom"`
	Timeto   int64 `gorm:"not null" json:"timeto"`
	Weekday  int64 `gorm:"not null" json:"weekday"`
}
