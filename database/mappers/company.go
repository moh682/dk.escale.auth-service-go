package mappers

import (
	"dk.escale.auth-service/database"
	"dk.escale.auth-service/database/models"

	uuid "github.com/satori/go.uuid"
)

// CreateCompany creates a company.
func CreateCompany(company *models.Company) error {
	db := database.GetConnection()
	tx := db.Create(company)
	if tx.Error != nil {
		return tx.Error
	}
	return tx.Error
}

// FindCompanyByCvr Finds a Company by cvr.
func FindCompanyByCvr(cvr int) (*models.Company, error) {
	db := database.GetConnection()
	company := &models.Company{
		Cvr: cvr,
	}
	tx := db.Find(company)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return company, nil
}

// FindAllEmployeesOfCompanyByUID Returs all company employees based on cvr
func FindAllEmployeesOfCompanyByUID(uid uuid.UUID) ([]models.Employee, error) {
	db := database.GetConnection()
	var employees []models.Employee
	tx := db.Where(&models.Employee{CompanyID: uid}).Find(&employees)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return employees, nil
}
