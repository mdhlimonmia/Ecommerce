package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//start time
		start := time.Now()

		//call the next handler
		fmt.Println("hit Logger")
		next.ServeHTTP(w, r)

		//end time
		duration := time.Since(start)

		//log the request method, path and duration
		tm := r.Method + " " + r.URL.Path + " " + duration.String()
		log.Println(r.URL.Path)
		log.Println(tm)
	})
}
