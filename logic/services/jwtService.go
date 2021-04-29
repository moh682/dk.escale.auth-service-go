package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"dk.escale.auth-service/config"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// TokenPayload token payload
type TokenPayload struct {
	jwt.StandardClaims
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	CompanyName string    `json:"companyName"`
	Role        string    `json:"role"`
	UserID      uuid.UUID `json:"userId"`
}

var env = config.GetEnvironments()

// SignToken Creates a token
func SignToken(
	userID uuid.UUID,
	userEmail string,
	userName string,
	userRole string,
	CompanyName string,
) string {
	payload := TokenPayload{
		Email:       userEmail,
		Name:        userName,
		UserID:      userID,
		Role:        userRole,
		CompanyName: CompanyName,
		StandardClaims: jwt.StandardClaims{
			Audience:  "dk.escale.*",
			ExpiresAt: time.Now().Unix() + (1000 * 60 * 60 * 24),
			Issuer:    env.JWTIssuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(env.JWTSecretKey))
	if err != nil {
		log.Fatal(err)
	}
	return signedToken
}

// GetTokenPayload returns the token payload
func GetTokenPayload(tokenString string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenPayload{}, tokenParseErrorHandler)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid Signing")
	}
	payload, ok := token.Claims.(*TokenPayload)
	if !ok {
		return nil, errors.New("Could not parse payload")
	}
	err = validateToken(payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func tokenParseErrorHandler(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	if token.Header["alg"] != "HS256" {
		return nil, fmt.Errorf("Incorrect encryption algorithm: %v", token.Header["alg"])
	}

	return []byte(env.JWTSecretKey), nil
}

func validateToken(payload *TokenPayload) error {
	err := payload.Valid()
	iss := config.GetEnvironments().JWTIssuer
	if err != nil {
		return err
	}
	if payload.IssuedAt <= 0 {
		return errors.New("Invalid iss")
	}
	if payload.Issuer != iss {
		return errors.New("Invalid Issuer")
	}
	return nil
}
