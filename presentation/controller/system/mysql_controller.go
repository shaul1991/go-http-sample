package controller

import (
	"encoding/json"
	"net/http"

	"go-http/core/database/mysql"
)

func MySQLHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]interface{})

	err := mysql.Ping()
	if err != nil {
		response["status"] = "error"
		response["message"] = "MySQL connection failed"
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		response["status"] = "success"
		response["message"] = "MySQL connection is healthy"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}
