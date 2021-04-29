package routes

import (
	"fmt"
)

var (
	// BaseAPI /api
	BaseAPI = "/api"
	// RegistationHandlerEndpoint /auth/registration
	RegistationHandlerEndpoint = fmt.Sprintf("%s/auth/registration", BaseAPI)
	// AddCompanyEndpoint /companies
	AddCompanyEndpoint = fmt.Sprintf("%s/companies", BaseAPI)
	// GetAllCompaniesEndpoint /companies/all
	GetAllCompaniesEndpoint = fmt.Sprintf("%s/companies/all", BaseAPI)
	// GetAllEmployeesEndpoint /employees
	GetAllEmployeesEndpoint = fmt.Sprintf("%s/employees", BaseAPI)
	// GetDocsEndpoint /docs
	GetDocsEndpoint = fmt.Sprintf("%s/docs", BaseAPI)
)
