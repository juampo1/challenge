package auth

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ValidateUser checks for a token and validates it
// before allowing the method to execute
func ValidateUser(key []byte, r http.Request) bool {
	reqToken := r.Header.Get("Authorization")

	if reqToken == "" {
		return false
	}

	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
