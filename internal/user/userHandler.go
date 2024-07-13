package user

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login handler called")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body")
		w.WriteHeader(400)
		return
	}
	var loginData map[string]string
	err = json.Unmarshal(body, &loginData)
	if err != nil {
		fmt.Println("Error in parsing request body")
		w.WriteHeader(400)
		return
	}
	username := loginData["username"]
	password := loginData["password"]

	_, err = io.WriteString(w, fmt.Sprintf("Welcome to login, username: %s, password: %s\n", username, password))
	if err != nil {
		w.WriteHeader(500)
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
