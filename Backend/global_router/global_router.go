package global_router

import (
	"ecommerce/handle_cors"
	"fmt"
	"net/http"
)

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleRequest := func(w http.ResponseWriter, r *http.Request) {
		handle_cors.HandleCors(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			fmt.Println("OPTIONS hit............")
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleRequest)
}
