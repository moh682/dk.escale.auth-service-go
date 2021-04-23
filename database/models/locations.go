package models

// Location model
type Location struct {
	Base
	City    string `gorm:"not null"`
	Address string `gorm:"not null"`
	Country string `gorm:"not null"`
	ZipCode int    `gorm:"not null"`
}
