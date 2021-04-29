package routes

import (
	"fmt"
	"net/http"
)

// GetAllEmployees Retrieve all users
func GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {

	// conn := db.OpenConnection()

	fmt.Fprintf(w, "GetAllEmployees hit")
}

// GetAllEmployeesInDepartment Retrieve all users
func GetAllEmployeesInDepartmentHandler(w http.ResponseWriter, r *http.Request) {

	// conn := db.OpenConnection()

	fmt.Fprintf(w, "GetAllEmployeesInDepartment hit")
}

// AddEmployee Retrieve all users
func AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	// conn := db.OpenConnection()

	fmt.Fprintf(w, "AddEmployee hit")
}
