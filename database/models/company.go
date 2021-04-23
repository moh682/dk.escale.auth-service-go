package models

import (
	"gorm.io/gorm"
)

// Company model
type Company struct {
	Base
	Departments []Department
	Cvr         int    `gorm:"not null;unique" json:"cvr"`
	Name        string `gorm:"not null" json:"name"`
}

// BeforeCreate creates
func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}
