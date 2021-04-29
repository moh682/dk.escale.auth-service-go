package routes

import (
	"fmt"
	"net/http"

	"dk.escale.auth-service/presentation/middlewares"

	"dk.escale.auth-service/config"
	"github.com/gorilla/mux"
)

// GetMainAPI generate a main api
func GetMainAPI() *mux.Router {
	server := mux.NewRouter()
	env := config.GetEnvironments()

	if env == nil {
		fmt.Println("Environments are nil")
	}

	protectedRoutes := server.Methods(http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPut).Subrouter()
	protectedRoutes.HandleFunc(GetAllEmployeesEndpoint, GetAllEmployeesHandler).Methods("GET")
	protectedRoutes.HandleFunc(GetAllCompaniesEndpoint, GetAllCompaniesHandler).Methods("GET")
	protectedRoutes.HandleFunc(AddCompanyEndpoint, AddCompanyHandler).Methods("POST")
	protectedRoutes.Use(middlewares.AuthorizationMiddleware)

	server.HandleFunc(GetDocsEndpoint, GetDocsHandler).Methods("GET")
	server.HandleFunc(RegistationHandlerEndpoint, RegistrationHandler).Methods("POST")

	return server
}
