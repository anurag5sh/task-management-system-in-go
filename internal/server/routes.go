package server

import (
	"net/http"

	"task-management-system/internal/user"
)

func CreateRoutes() {
	createUserRoutes()
}

func createUserRoutes() {
	http.HandleFunc("/login", user.LoginHandler)
}
