package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"task-management-system/internal/database"
	dbQuery "task-management-system/internal/database/sqlc/sqlc-autogen"
)

type ReqHandler func(w http.ResponseWriter, r *http.Request)

func LoginHandler(db *database.Database) ReqHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login handler called")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body")
			w.WriteHeader(400)
			return
		}
		var loginData map[string]string

		if err = json.Unmarshal(body, &loginData); err != nil {
			fmt.Println("Error in parsing request body")
			w.WriteHeader(400)
			return
		}
		username := loginData["username"]
		password := loginData["password"]

		queries := dbQuery.New(db.Db)
		userPasswordHash, err := queries.GetUserPassword(db.Ctx, username)

		if err = bcrypt.CompareHashAndPassword([]byte(userPasswordHash), []byte(password)); err != nil {
			w.WriteHeader(404)
			_, err := io.WriteString(w, "Wrong username or password")
			if err != nil {
				w.WriteHeader(500)
			}

			return
		}

		var (
			key []byte
			t   *jwt.Token
			s   string
		)

		key = []byte("mykey")
		t = jwt.New(jwt.SigningMethodHS256)
		s, err = t.SignedString(key)
		if err != nil {
			fmt.Println("Unable to sign JWT")
		}

		w.Header().Set("Authorization", "Bearer "+s)

		if _, err = io.WriteString(w, fmt.Sprintf("Logged in successfully")); err != nil {
			fmt.Printf("\nUnable to send response %s", err)
			w.WriteHeader(500)
			return
		}
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error in hashing password")
		return "", err
	}

	return string(hash), nil
}

func RegisterHandler(db *database.Database) ReqHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		userData := make(map[string]string)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error in reading request body")
			w.WriteHeader(400)
			return
		}
		if err = json.Unmarshal(body, &userData); err != nil {
			fmt.Println("Error in parsing request body")
			w.WriteHeader(400)
			return
		}
		if _, ok := userData["username"]; !ok {
			fmt.Println("Username not found in request")
			w.WriteHeader(400)
			return
		}
		if _, ok := userData["password"]; !ok {
			fmt.Println("Password not found in request")
			w.WriteHeader(400)
			return
		}
		if _, ok := userData["email"]; !ok {
			fmt.Println("Email not found in request")
			w.WriteHeader(400)
			return
		}

		passwordHash, err := hashPassword(userData["password"])
		if err != nil {
			fmt.Println("Error in hashing password")
			w.WriteHeader(500)
			return
		}

		queries := dbQuery.New(db.Db)
		err = queries.CreateUser(db.Ctx, dbQuery.CreateUserParams{Username: userData["username"],
			PasswordHash: passwordHash, Email: sql.NullString{String: userData["email"], Valid: true}, CreatedAt: time.Now()})
		if err != nil {
			w.WriteHeader(400)
			if _, err := io.WriteString(w, fmt.Sprintf("Error in creating user: %s", err)); err != nil {
				fmt.Printf("error in sending response: %s", err)
			}
			return
		}

		_, err = io.WriteString(w, "User successfully created")
		if err != nil {
			fmt.Println("Error in writing response")
			w.WriteHeader(500)
			return
		}
	}
}
