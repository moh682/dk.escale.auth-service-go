package mappers

import (
	"dk.escale.auth-service/database"
	"dk.escale.auth-service/database/models"
	uuid "github.com/satori/go.uuid"
)

// CreateEmployee Needs companyID.
func CreateEmployee(email string, password string, companyID uuid.UUID) error {
	db := database.GetConnection()
	employee := models.Employee{
		Email:     email,
		Password:  password,
		CompanyID: companyID,
	}
	tx := db.Create(&employee)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
