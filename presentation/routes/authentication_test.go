package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"dk.escale.auth-service/presentation/routes"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationHandler(t *testing.T) {
	request, _ := http.NewRequest("POST", routes.RegistationHandlerEndpoint, nil)
	response := httptest.NewRecorder()
	routes.GetMainAPI().ServeHTTP(response, request)
	assert.EqualValues(t, 200, response.Code, "Expected response code to equal 200")
}
