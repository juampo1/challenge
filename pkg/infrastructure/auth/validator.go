package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ValidateUser checks for a token and validates it
// before allowing the method to execute
func ValidateUser(next http.HandlerFunc, key []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")

		if reqToken == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
				return key, nil
			})

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			if !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			ctx := context.WithValue(r.Context(), "Token", token)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
