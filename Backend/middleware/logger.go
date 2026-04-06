package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//start time
		start := time.Now()

		//call the next handler
		next.ServeHTTP(w, r)

		//end time
		duration := time.Since(start)

		//log the request method, path and duration
		tm := r.Method + " " + r.URL.Path + " " + duration.String()
		log.Println(tm)
	}
}
