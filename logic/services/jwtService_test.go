package services

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"
)

func TestSignToken(t *testing.T) {
	token := SignToken(uuid.NewV4(), "email", "George", "escale", "admin")
	assert.NotNil(t, token, "Token is empty")
}

func TestGetTokenPayloadForValidToken(t *testing.T) {
	name := "George"
	email := "email"
	companyName := "escale"
	token := SignToken(uuid.NewV4(), email, name, companyName, "admin")
	payload, err := GetTokenPayload(token)
	assert.NotNil(t, payload, "Expect Token To Not be Nil")
	assert.Nil(t, err, "Expect Error To be nil")
	assert.Equal(t, name, payload.Name, fmt.Sprintf("Expected name to be %v", name))
	assert.Equal(t, email, payload.Email, fmt.Sprintf("Expected email to be %v", email))
}

// Security Check for known JWT security breaches
// source: https://apisecurity.io/issue-56-common-jwt-attacks-owasp-api-security-top-10-cheatsheet/
// TODO Test for all security flaws
func TestTokenValidationWithDifferentAlgorithmInHeader(t *testing.T) {
	// Token has algorithm "none"
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJpZCI6IjEzMzciLCJ1c2VybmFtZSI6ImJpem9uZSIsImlhdCI6MTU5NDIwOTYwMCwicm9sZSI6ImFkbWluIn0.iBR-xKQpMntc8lc4lzvCYyVKSgIU93BK6g83vufWfuE"
	payload, err := GetTokenPayload(token)
	assert.NotNil(t, err, "Expect Token To Not be Nil")
	assert.Nil(t, payload, "Expect Error To be nil")
}
