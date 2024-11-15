package route

import (
	mainController "go-http/internal/domain/main/controller"
	systemController "go-http/internal/domain/system/controller"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	systemController.HealthCheck(w, r)
}

func MongoDBHandler(w http.ResponseWriter, r *http.Request) {
	systemController.MongoDBHealthCheck(w, r)
}

func MySQLHandler(w http.ResponseWriter, r *http.Request) {
	systemController.MySQLHealthCheck(w, r)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	systemController.DashboardHandler(w, r)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	mainController.MainResponse(w, r)
}
