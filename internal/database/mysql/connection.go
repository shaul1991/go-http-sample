package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenConns    = 100
	maxIdleConns    = 10
	connMaxLifetime = 30 * time.Minute
	connMaxIdleTime = 5 * time.Minute
)

// DB is a MySQL database instance
var DB *sql.DB

// Connect establishes a connection to MySQL with connection pool
func Connect(user, password, host, port, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	// Verify connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	DB = db
	log.Printf("Connected to MySQL! Max connections: %d, Idle connections: %d",
		maxOpenConns, maxIdleConns)
	return nil
}

// Disconnect closes the MySQL connection pool
func Disconnect() error {
	if DB == nil {
		return nil
	}

	err := DB.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %v", err)
	}

	log.Println("Disconnected from MySQL")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
}

// Ping checks if the database connection is alive
func Ping() error {
	if DB == nil {
		return fmt.Errorf("database connection not initialized")
	}
	return DB.Ping()
} 