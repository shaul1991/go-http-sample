package controller

import (
	"fmt"
	"net/http"
)

// HealthCheck responds with a 200 status and "ok" message
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
