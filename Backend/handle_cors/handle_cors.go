package handle_cors

import "net/http"

func HandleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //Alow all origin by "*"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
