package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dk.escale.auth-service/database"
	"dk.escale.auth-service/database/models"
)

type companyDTO struct {
	name string
}

// GetAllCompanies get all companies
func GetAllCompanies(w http.ResponseWriter, r *http.Request) {

	// conn := db.OpenConnection()

	fmt.Fprintf(w, "GetAllEmployees hit")
}

// GetCompany get a single company
func GetCompany(w http.ResponseWriter, r *http.Request) {

	// conn := db.OpenConnection()

	fmt.Fprintf(w, "GetAllEmployeesInDepartment hit")
}

// AddCompany create a company
func AddCompany(w http.ResponseWriter, r *http.Request) {
	// conn := db.OpenConnection()

	var company models.Company

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&company)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	DB := database.GetConnection()

	DB.Create(&models.Company{
		Cvr:  company.Cvr,
		Name: company.Name,
		// Departments: []models.DepartmentModel{
		// 	{
		// 		Name: "new",
		// 	},
		// },
	})

	// DB.Find(company, company)

	w.Header().Set("Content-Type", "Application/json")
	fmt.Fprintf(w, "Company: %+v", company)
}
