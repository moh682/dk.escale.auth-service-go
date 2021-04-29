package mappers_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"dk.escale.auth-service/database/mappers"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	company, err := mappers.FindCompanyByCvr(CompanyCVR)
	assert.Nil(t, err, "Expected to find company without errors")
	err = mappers.CreateEmployee("newEmail@gmail.com", "password", company.ID, uuid.Nil)
	assert.Nil(t, err, "Expected to add employee without errors")
}
