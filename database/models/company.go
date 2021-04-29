package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Company model
type Company struct {
	Base
	Employees  []Employee
	LocationID uuid.UUID `json:"location_id"`
	Location   Location
	Cvr        int    `gorm:"not null;unique" json:"cvr"`
	Name       string `gorm:"not null" json:"name"`
	Email      string
}

// BeforeCreate creates
func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}
