package user

import (
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	//Decode the JSON request body into a ReqLogin struct
	var TryUser ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&TryUser)
	if err != nil {
		log.Println(err)
		util.SendData(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	//Check if the email exists in the database
	findEmail := h.userRepo.FindUser(TryUser.Email)
	if findEmail == false {
		util.SendData(w, "Email Not Found.", http.StatusBadRequest)
		return
	}
	//Authenticate the user with the provided email and password
	user, err := h.userRepo.AuthUser(TryUser.Email, TryUser.Password)
	if err != nil {
		util.SendData(w, "Error authenticating user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		util.SendData(w, "Invalid Password.", http.StatusBadRequest)
		return
	}

	//Create a JWT token for the authenticated user
	id := strconv.Itoa(user.ID)
	accessToken, err := util.CreateJwt(util.Payload{
		Sub:         id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	}, h.cnf.JwtSecretKey)
	if err != nil {
		util.SendData(w, "Error creating JWT token", http.StatusInternalServerError)
		return
	}

	//Send the JWT token as a response to the client
	util.SendData(w, accessToken, http.StatusCreated)
}
