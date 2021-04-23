package mappers

import (
	"dk.escale.auth-service/database"
	"dk.escale.auth-service/database/models"
)

// CreateCompany creates a company
func CreateCompany(company *models.Company) error {
	db := database.GetConnection()

	tx := db.Create(company)

	if tx.Error != nil {
		return tx.Error
	}

	// tx := db.Create(location)

	// if tx.Error != nil {
	// 	return tx.Error
	// }

	return tx.Error
}

// // CreateCompany creates a company
// func CreateCompany(company *models.CompanyModel) {
// 	db := database.GetConnection()
// 	companyID := uuid.NewV4().String()

// 	db.Create(&models.CompanyModel{
// 		Base: models.Base{
// 			ID: companyID,
// 		},
// 		Cvr:  company.Cvr,
// 		Name: company.Name,
// 		Departments: []models.DepartmentModel{
// 			{
// 				CompanyID: companyID,
// 				Name:      "new",
// 			},
// 		},
// 	})

// }
