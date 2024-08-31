package middleware

import (
	"net/http"
)

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unauthorize", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
