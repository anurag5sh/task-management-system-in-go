package task

import (
	"encoding/json"
	"fmt"
	"net/http"

	"task-management-system/internal/database"
	sqlc "task-management-system/internal/database/sqlc/sqlc-autogen"
	"task-management-system/internal/model"
)

func GetAllTasks(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := r.Context().Value("claims").(*model.UserClaims)

		if session.UserId != 0 {
			fmt.Println("Userdata not found in session")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		queries := sqlc.New(db.Db)

		tasks, err := queries.GetAllTasks(db.Ctx, session.UserId)
		if err != nil {
			fmt.Printf("Error getting tasks: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(tasks)
	}
}

func GetTask(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func CreateTask(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func UpdateTask(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
