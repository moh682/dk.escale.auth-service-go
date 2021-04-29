package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Employee model
type Employee struct {
	Base
	// RoleID       string
	CompanyID uuid.UUID `gorm:"type:uuid;foreignKey:id"`
	// DepartmentID uuid.UUID `gorm:"type:uuid"`
	Role        string `gorm:"not null;default:employee"`
	Name        string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	MobilePhone string `gorm:"not null"`
}

// BeforeCreate creates
func (u *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	if u.CompanyID == uuid.Nil {
		// var department Department = Department{
		// 	Base: Base{
		// 		ID: u.DepartmentID,
		// 	},
		// }
		// tx.First(&department)
		// u.CompanyID = department.CompanyID
		// u.DepartmentID = department.ID
		if tx.Error != nil {
			return tx.Error
		}
		password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return err
		}
		u.Password = string(password)
	}
	return nil
}
