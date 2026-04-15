package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Create_user(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	createUser := newUser.Store()

	util.SendData(w, createUser, http.StatusCreated)
}
