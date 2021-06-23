package auth

import "net/http"

// ValidateUser checks for a token and validates it
// before allowing the method to execute
func ValidateUser(_ http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate token
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
}
