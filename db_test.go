package main

import (
	"math/rand"
	"testing"

	"dk.escale.auth-service/database"
	"dk.escale.auth-service/database/mappers"

	"dk.escale.auth-service/database/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var _ = Describe("Test ConnectionService", func() {

	// ************************ Connection ************************
	Describe("Open Connection To Postgres", func() {
		Context("Open Connection To Postgres", func() {
			It("Expect Connection to be open", func() {
				DB := database.GetConnection()
				Expect(DB.Error).To(BeNil())
			})
		})
	})

	// ************************ Company Test ************************
	Describe("Create a Company", func() {
		Context("Create a Company", func() {
			It("ECreate a Company", func() {
				company := &models.Company{
					Cvr:  rand.Intn(8),
					Name: "escale",
					Departments: []models.Department{models.Department{
						Name: "HeadDepartment",
						Location: models.Location{
							Address: "Some nice street",
							City:    "Nice city",
							Country: "Nice Country",
							ZipCode: 21112,
						},
						Employees: []models.Employee{
							{
								Name:        "welcome",
								Password:    "Password",
								Email:       "email@email.com",
								MobilePhone: "123123123",
							},
						},
					}},
				}
				err := mappers.CreateCompany(company)
				Expect(err).To(BeNil())
			})
		})
	})

	// Describe("Create a Company With Department", func() {
	// 	Context("Open Connection To Postgres", func() {
	// 		It("Expect Connection to be open", func() {
	// 			company := &models.CompanyModel{
	// 				Cvr:  12333,
	// 				Name: "escale",
	// 				Departments: []models.DepartmentModel{
	// 					{
	// 						Name: "Head Department",
	// 					},
	// 				},
	// 			}
	// 			mappers.CreateCompany(company)
	// 		})
	// 	})
	// })

})
