package controller

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go-http/internal/database/mongodb" 
	"go-http/internal/database/mysql"
)

type DatabaseStatus struct {
	Name      string                 `json:"name"`
	Status    string                 `json:"status"`
	Message   string                 `json:"message"`
	Stats     map[string]interface{} `json:"stats"`
	PoolStats map[string]interface{} `json:"poolStats"`
}

type DashboardData struct {
	Databases []DatabaseStatus
	Timestamp string
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/json" {
		handleJSONResponse(w)
		return
	}
	handleHTMLResponse(w)
}

func handleJSONResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	statuses := checkDatabaseStatuses()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"databases": statuses,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func handleHTMLResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")

	// 템플릿 파일 경로
	tmplPath := filepath.Join("internal", "domain", "system", "templates", "dashboard.html")
	
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := DashboardData{
		Databases: checkDatabaseStatuses(),
		Timestamp: time.Now().Format(time.RFC3339),
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

func checkDatabaseStatuses() []DatabaseStatus {
	var statuses []DatabaseStatus

	// Check MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	mongoStatus := DatabaseStatus{
		Name:      "MongoDB",
		Stats:     make(map[string]interface{}),
		PoolStats: make(map[string]interface{}),
	}

	if err := mongodb.Client.Ping(ctx, nil); err != nil {
		mongoStatus.Status = "error"
		mongoStatus.Message = "Connection failed"
	} else {
		mongoStatus.Status = "success"
		mongoStatus.Message = "Connected successfully"

		// Get MongoDB server status
		serverStatus := bson.M{}
		err := mongodb.Client.Database("admin").RunCommand(ctx, bson.D{{Key: "serverStatus", Value: 1}}).Decode(&serverStatus)
		if err == nil {
			if connections, ok := serverStatus["connections"].(bson.M); ok {
				mongoStatus.Stats["Current Connections"] = connections["current"]
				mongoStatus.Stats["Available Connections"] = connections["available"]
			}
			if mem, ok := serverStatus["mem"].(bson.M); ok {
				mongoStatus.Stats["Memory Resident (MB)"] = mem["resident"]
				mongoStatus.Stats["Memory Virtual (MB)"] = mem["virtual"]
			}
		}

		// Get MongoDB pool stats
		poolStats := mongodb.Client.NumberSessionsInProgress()
		mongoStatus.PoolStats["Active Sessions"] = poolStats
		mongoStatus.PoolStats["Min Pool Size"] = mongodb.GetMinPoolSize()
		mongoStatus.PoolStats["Max Pool Size"] = mongodb.GetMaxPoolSize()
	}
	statuses = append(statuses, mongoStatus)

	// Check MySQL
	mysqlStatus := DatabaseStatus{
		Name:      "MySQL",
		Stats:     make(map[string]interface{}),
		PoolStats: make(map[string]interface{}),
	}

	if err := mysql.Ping(); err != nil {
		mysqlStatus.Status = "error"
		mysqlStatus.Message = "Connection failed"
	} else {
		mysqlStatus.Status = "success"
		mysqlStatus.Message = "Connected successfully"

		// Get MySQL status variables
		db := mysql.GetDB()
		rows, err := db.Query("SHOW GLOBAL STATUS WHERE Variable_name IN ('Threads_connected', 'Max_used_connections', 'Threads_running', 'Bytes_received', 'Bytes_sent')")
		if err == nil {
			defer rows.Close()
			
			for rows.Next() {
				var name string
				var value string
				if err := rows.Scan(&name, &value); err == nil {
					switch name {
					case "Threads_connected":
						mysqlStatus.Stats["Active Connections"] = value
					case "Max_used_connections":
						mysqlStatus.Stats["Max Used Connections"] = value
					case "Threads_running":
						mysqlStatus.Stats["Running Threads"] = value
					case "Bytes_received":
						mysqlStatus.Stats["Bytes Received"] = value
					case "Bytes_sent":
						mysqlStatus.Stats["Bytes Sent"] = value
					}
				}
			}
		}

		// Get MySQL pool stats
		stats := db.Stats()
		mysqlStatus.PoolStats["Open Connections"] = stats.OpenConnections
		mysqlStatus.PoolStats["In Use Connections"] = stats.InUse
		mysqlStatus.PoolStats["Idle Connections"] = stats.Idle
		mysqlStatus.PoolStats["Max Open Connections"] = stats.MaxOpenConnections
	}
	statuses = append(statuses, mysqlStatus)

	return statuses
} 