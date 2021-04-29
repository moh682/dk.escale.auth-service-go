package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"dk.escale.auth-service/logic/services"
)

// MiddlewareErrorResponse Response type for invalid AuthorizationMiddleware
type MiddlewareErrorResponse struct {
	error string
}

// AuthorizationMiddleware Middleware to authenticate routes.
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// Check for correct token

		trimmedToken := strings.ReplaceAll(r.Header.Get("Authorization"), " ", "")
		token := strings.Replace(trimmedToken, "Bearer", "", 1)

		if token == "" {
			w.WriteHeader(403)
			return
		}

		_, err := services.GetTokenPayload(token)

		if err != nil {
			w.WriteHeader(401)
			w.Header().Set("Content-Type", "application/json")
			body, _ := json.Marshal(&MiddlewareErrorResponse{
				error: fmt.Sprintf("Authorization failed, token parse error: %v", err.Error()),
			})
			w.Write([]byte(body))
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
