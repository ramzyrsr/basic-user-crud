package middleware

import (
	"net/http"

	"lemonilo.app/auth"
	response "lemonilo.app/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		next(w, r)
	}
}
