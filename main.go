package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"dk.escale.auth-service/config"
	"dk.escale.auth-service/presentation/routes"
)

func main() {
	env := config.GetEnvironments()
	if env == nil {
		fmt.Println("Environments are nil")
	}
	r := mux.NewRouter()

	r.HandleFunc("/employees", routes.GetAllEmployees).Methods("GET")
	r.HandleFunc("/companies/all", routes.GetAllCompanies).Methods("GET")
	r.HandleFunc("/companies", routes.AddCompany).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), r)
}
