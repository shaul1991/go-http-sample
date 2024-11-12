package route

import (
	"net/http"
	"go-http/internal/domain/system/controller"
)

// HealthHandler handles the /health route
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	controller.HealthCheck(w, r)
}

// RootHandler handles the / route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	controller.ReturnMessage(w, r)
}
