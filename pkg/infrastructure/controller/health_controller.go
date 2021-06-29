package controller

import (
	"encoding/json"
	"net/http"

	"github.com/challenge/pkg/domain"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(domain.Health{Health: "ok"})
	}
}
