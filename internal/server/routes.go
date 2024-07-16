package server

import (
	"net/http"
	"task-management-system/internal/database"

	"task-management-system/internal/user"
)

func CreateRoutes(db *database.Database) {
	createUserRoutes(db)
}

func createUserRoutes(db *database.Database) {
	http.HandleFunc("/login", user.LoginHandler(db))
}
