package middleware

import (
	"fmt"
	"net/http"
)

func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			fmt.Println("OPTIONS hit............")
			return
		}
		fmt.Println("hit Preflight")
		next.ServeHTTP(w, r)
	})
}
