package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go-http/internal/database/mongodb"
)

func MongoDBHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// MongoDB 연결 상태 확인
	err := mongodb.Client.Ping(ctx, nil)

	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]interface{})

	if err != nil {
		response["status"] = "error"
		response["message"] = "MongoDB connection failed"
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		response["status"] = "success"
		response["message"] = "MongoDB connection is healthy"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
} 