package mappers_test

import (
	"math/rand"
	"testing"
	"time"

	"dk.escale.auth-service/database/mappers"
	"dk.escale.auth-service/database/models"
	"github.com/stretchr/testify/assert"
)

var CompanyCVR = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000)

func TestAddCompanyMapper(t *testing.T) {
	employee := models.Employee{
		Name:        "welcome",
		Password:    "Password",
		Email:       "email@email.com",
		MobilePhone: "123123123",
	}
	location := models.Location{
		Address: "Some nice street",
		City:    "Nice city",
		Country: "Nice Country",
		ZipCode: 21112,
	}
	// department := models.Department{
	// 	Name:      "HeadDepartment",
	// 	Location:  location,
	// 	Employees: []models.Employee{employee},
	// }
	company := &models.Company{
		Cvr:       CompanyCVR,
		Name:      "escale",
		Employees: []models.Employee{employee},
		locationID := 
	}
	err := mappers.CreateCompany(company)
	assert.Nil(t, err, "Expected to be no errors")
}

func TestFindCompanyByCvr(t *testing.T) {
	company, err := mappers.FindCompanyByCvr(CompanyCVR)
	assert.Nil(t, err, "Expected to be no errors")
	assert.NotNil(t, company.ID, "Expected companyID to not be nil")
}

func TestFindAllEmplyeesByCompanyUid(t *testing.T) {
	company, err := mappers.FindCompanyByCvr(CompanyCVR)
	employees, err := mappers.FindAllEmployeesOfCompanyByUID(company.ID)
	assert.Equal(t, 1, len(employees), "Expected the length of employees to be 1")
	assert.Equal(t, company.ID, employees[0].CompanyID, "Expected employee.CompanyID to be equal company.ID")
	assert.Nil(t, err, "Expected to be no errors")
	assert.NotNil(t, company.ID, "Expected companyID to not be nil")
}
