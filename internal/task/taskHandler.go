package task

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/anurag5sh/task-management-system-in-go/internal/database"
	sqlc "github.com/anurag5sh/task-management-system-in-go/internal/database/sqlc/sqlc-autogen"
	"github.com/anurag5sh/task-management-system-in-go/internal/model"
)

func GetAllTasks(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := r.Context().Value("claims").(*model.UserClaims)

		if session.Uid == 0 {
			fmt.Println("Userdata not found in session")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		queries := sqlc.New(db.Db)

		tasks, err := queries.GetAllTasks(db.Ctx, session.Uid)
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
		session, _ := r.Context().Value("claims").(*model.UserClaims)

		if session.Uid == 0 {
			fmt.Println("Userdata not found in session")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		idValue := r.PathValue("id")
		id, err := strconv.ParseInt(idValue, 0, 64)
		queries := sqlc.New(db.Db)

		task, err := queries.GetTask(db.Ctx, sqlc.GetTaskParams{
			ID:     id,
			UserID: session.Uid,
		})

		if err != nil {
			if task.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("Task not found"))
				return
			}
			fmt.Printf("Error getting task: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(task)
	}
}

func CreateTask(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := r.Context().Value("claims").(*model.UserClaims)

		if session.Uid == 0 {
			fmt.Println("Userdata not found in session")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body")
			w.WriteHeader(400)
			return
		}
		var reqData map[string]string

		if err = json.Unmarshal(body, &reqData); err != nil {
			fmt.Println("Error in parsing request body")
			w.WriteHeader(400)
			return
		}
		queries := sqlc.New(db.Db)
		err = queries.CreateTask(db.Ctx, sqlc.CreateTaskParams{
			Title:       reqData["title"],
			Description: reqData["description"],
			Status:      reqData["status"],
			UserID:      session.Uid,
			CreatedAt:   time.Now(),
			UpdatedAt:   sql.NullTime{},
		})
		if err != nil {
			fmt.Printf("Error creating task: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("Task created successfully"))
	}
}
func UpdateTask(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := r.Context().Value("claims").(*model.UserClaims)

		if session.Uid == 0 {
			fmt.Println("Userdata not found in session")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body")
			w.WriteHeader(400)
			return
		}
		var reqData map[string]string

		if err = json.Unmarshal(body, &reqData); err != nil {
			fmt.Println("Error in parsing request body")
			w.WriteHeader(400)
			return
		}

		idValue := r.PathValue("id")
		taskId, err := strconv.ParseInt(idValue, 0, 64)
		if err != nil {
			w.WriteHeader(400)
			_, _ = io.WriteString(w, "invalid task ID")
			return
		}

		queries := sqlc.New(db.Db)

		if exist, err := queries.IsTaskExist(db.Ctx, taskId); exist != 1 {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(fmt.Sprintf("Error updating task: %v\n", err)))
				return

			}
			w.WriteHeader(400)
			_, _ = w.Write([]byte("Task doesn't exist"))
			return
		}

		err = queries.UpdateTask(db.Ctx, sqlc.UpdateTaskParams{
			ID:          taskId,
			Title:       reqData["title"],
			Description: reqData["description"],
			Status:      reqData["status"],
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})
		if err != nil {
			fmt.Printf("Error updating task: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Task updated successfully"))
	}
}
