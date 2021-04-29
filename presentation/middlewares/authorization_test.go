package middlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"dk.escale.auth-service/presentation/routes"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizationMiddlewareWithNoToken(t *testing.T) {
	request, _ := http.NewRequest("GET", routes.GetAllCompaniesEndpoint, nil)
	request.Header.Set("Authorization", "Bearer ")
	response := httptest.NewRecorder()
	routes.GetMainAPI().ServeHTTP(response, request)
	assert.EqualValues(t, 403, response.Code, "Expected response code to equal 403")
}

func TestAuthorizationMiddlewareWithWrongTokenAlgorithm(t *testing.T) {
	request, _ := http.NewRequest("GET", routes.GetAllCompaniesEndpoint, nil)
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJpZCI6IjEzMzciLCJ1c2VybmFtZSI6ImJpem9uZSIsImlhdCI6MTU5NDIwOTYwMCwicm9sZSI6ImFkbWluIn0"
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	response := httptest.NewRecorder()
	routes.GetMainAPI().ServeHTTP(response, request)
	assert.EqualValues(t, 401, response.Code, "Expected response code to equal 403")
}

// func TestAuthorizationMiddlewareForUnAuthorizedToken(t *testing.T) {
// 	// protectedEndpoint := routes.GetAllCompaniesEndpoint
// 	// request, _ := http.NewRequest(http.MethodGet, protectedEndpoint, nil)
// 	// response := httptest.NewRecorder()
// 	// GetMainAPI().ServeHTTP(response, request)
// 	// assert.EqualValues(t, 403, response.Code, "Expected response code to equal 403")
// }
