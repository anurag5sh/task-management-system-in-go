package server

import (
	"net/http"
	"github.com/anurag5sh/task-management-system-in-go/internal/database"
	"github.com/anurag5sh/task-management-system-in-go/internal/middleware"
	"github.com/anurag5sh/task-management-system-in-go/internal/task"
	"github.com/anurag5sh/task-management-system-in-go/internal/user"
)

func CreateRoutes(hs *HttpServer, db *database.Database) http.Handler {
	createUserRoutes(hs.Mux, db)
	createTaskRoutes(hs.Mux, db)
	return hs.Mux
}

func createUserRoutes(mux *http.ServeMux, db *database.Database) {
	mux.HandleFunc("POST /login", user.LoginHandler(db))
	mux.HandleFunc("POST /register", user.RegisterHandler(db))
}

func createTaskRoutes(mux *http.ServeMux, db *database.Database) {
	mux.HandleFunc("GET /tasks", middleware.CheckAuth(task.GetAllTasks(db)))
	mux.HandleFunc("POST /tasks", middleware.CheckAuth(task.CreateTask(db)))
	mux.HandleFunc("GET /tasks/{id}", middleware.CheckAuth(task.GetTask(db)))
	mux.HandleFunc("PUT /tasks/{id}", middleware.CheckAuth(task.UpdateTask(db)))
}
