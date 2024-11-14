package repository

import (
	"database/sql"
	"fmt"
	"go-http/core/database/mysql"
	"go-http/core/domain/user/model"
)

func CreateUser(name, email string) (string, error) {
	db := mysql.GetDB()
	result, err := db.Exec("INSERT INTO users (id, name, email) VALUES (UUID(), ?, ?)", name, email)
	if err != nil {
		return "", fmt.Errorf("could not insert user: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("could not retrieve last insert id: %v", err)
	}
	return fmt.Sprintf("%d", id), nil
}

func GetUserByID(id string) (*model.User, error) {
	db := mysql.GetDB()
	row := db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not query user: %v", err)
	}
	return &user, nil
}

func DeleteUser(id string) error {
	db := mysql.GetDB()
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}
	return nil
}
