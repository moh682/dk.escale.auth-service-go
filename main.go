package main

import (
	"fmt"
	"net/http"

	"dk.escale.auth-service/presentation/routes"

	"dk.escale.auth-service/config"
)

// Server application
var env = config.GetEnvironments()

func main() {
	http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), routes.GetMainAPI())
}
