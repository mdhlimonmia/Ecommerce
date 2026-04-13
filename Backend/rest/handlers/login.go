package handlers

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	//Decode the JSON request body into a ReqLogin struct
	var TryUser ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&TryUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	//Check if the email exists in the database
	findEmail := database.FindUser(TryUser.Email)
	if findEmail == false {
		http.Error(w, "Email Not Found.", http.StatusBadRequest)
		return
	}
	//Authenticate the user with the provided email and password
	user := database.AuthUser(TryUser.Email, TryUser.Password)
	if user == nil {
		http.Error(w, "Invalid Password.", http.StatusBadRequest)
		return
	}

	//Create a JWT token for the authenticated user
	cnf := config.GetConfig()
	id := strconv.Itoa(user.ID)
	accessToken, err := util.CreateJwt(util.Payload{
		Sub:         id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	}, cnf.JwtSecretKey)
	if err != nil {
		http.Error(w, "Error creating JWT token", http.StatusInternalServerError)
		return
	}

	//Send the JWT token as a response to the client
	util.SendData(w, accessToken, http.StatusCreated)
}
