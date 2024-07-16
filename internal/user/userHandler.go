package user

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"

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

		fmt.Println(s)
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
