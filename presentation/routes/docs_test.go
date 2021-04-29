package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocsHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", GetDocsEndpoint, nil)
	response := httptest.NewRecorder()
	GetMainAPI().ServeHTTP(response, request)
	assert.EqualValues(t, 200, response.Code, "Expected response code to equal 200")
}
