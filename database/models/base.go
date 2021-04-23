package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Base model
type Base struct {
	ID        uuid.UUID  `gorm:"type:UUID;" json:"_id"`
	CreatedAt time.Time  `gorm:"type:Date;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:Date;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:Date;" json:"deleted_at"`
}

// BeforeCreate creates
func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return nil
}

// BeforeSave creates
func (u *Base) BeforeSave(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return nil
}
