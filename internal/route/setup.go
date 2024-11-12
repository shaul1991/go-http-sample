package route

import (
	"net/http"
)

// SetupRoutes initializes and returns the HTTP handler with all routes configured
func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	
	setupSystemRoutes(mux)
	setupDefaultRoutes(mux)
	
	return mux
}

// setupSystemRoutes configures system-related routes like health check
func setupSystemRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/check/mongodb", MongoDBHandler)
}

// setupDefaultRoutes configures default routes
func setupDefaultRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", RootHandler)
}
