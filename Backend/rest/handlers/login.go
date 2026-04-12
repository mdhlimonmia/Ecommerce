package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"log"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var TryUser ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&TryUser)

	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	findEmail := database.FindUser(TryUser.Email)
	if findEmail == false {
		http.Error(w, "Email Not Found.", http.StatusBadRequest)
		return
	}

	user := database.AuthUser(TryUser.Email, TryUser.Password)

	if user == nil {
		http.Error(w, "Invalid Password.", http.StatusBadRequest)
		return
	}

	util.SendData(w, user, http.StatusCreated)
}
