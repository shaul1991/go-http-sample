package route

import (
	"net/http"

	mainController "go-http/presentation/controller/main"
	systemController "go-http/presentation/controller/system"
	userController "go-http/presentation/controller/user"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// System routes
	r.HandleFunc("/check/health", systemController.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/check/mongodb", systemController.MongoDBHealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/check/mysql", systemController.MySQLHealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/check/dashboard", systemController.DashboardHandler).Methods(http.MethodGet)

	// User routes
	r.HandleFunc("/api/users", userController.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/users", userController.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/users", userController.DeleteUserHandler).Methods(http.MethodDelete)

	// Default route
	r.HandleFunc("/", mainController.MainResponse).Methods(http.MethodGet)

	return r
}
