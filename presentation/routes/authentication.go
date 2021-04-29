package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dk.escale.auth-service/database/models"
)

// CompanyDTO for requests and responses
type CompanyDTO struct {
	Name  string
	Email string
	Cvr   int
}

// DepartmentDTO for requests and responses
type DepartmentDTO struct {
	Name  string
	Email string
	Cvr   int
}

type registrationResponse struct {
	Company *models.Company `json:"company"`
}
type registerRequest struct {
	Company    CompanyDTO
	Department DepartmentDTO
}

// RegistrationHandler Handler for registration
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	body, err := json.Marshal(&registrationResponse{
		Company: &models.Company{},
	})

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
		return
	}
	_, err = w.Write(body)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	return
}
