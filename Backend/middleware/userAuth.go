package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		pass := r.FormValue("password")

		if user != "admin" || pass != "1234" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
