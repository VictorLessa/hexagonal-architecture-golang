package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type ErrorMessage struct {
	Message string `json:"message"`
}


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authenticate")
		
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secureSecretText"), nil
		})

		if err != nil {
			response, _ := json.Marshal(ErrorMessage{
				Message: "Invalid token!",
			})


			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write(response)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}