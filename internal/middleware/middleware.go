package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"task-management-system/internal/model"
)

func CheckAuth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		if len(authHeader) < 1 {
			fmt.Println("Invalid Authorization header")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Invalid Authorization header"))
			return
		}

		jwtToken := strings.Trim(authHeader[1], " ")
		token, err := jwt.ParseWithClaims(jwtToken, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secret := os.Getenv("TMS_JWT_SECRET")
			if len(secret) == 0 {
				secret = "mykey"
			}
			return []byte(secret), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Invalid Authorization header"))
			return
		}

		if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Invalid Authorization header"))
			return
		}

	})
}
