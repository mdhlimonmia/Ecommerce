package user

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"log"
	"net/http"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) Create_user(w http.ResponseWriter, r *http.Request) {
	var newUser ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		log.Println(err)
		util.SendData(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" || newUser.Password == "" {
		util.SendData(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if h.svc.FindUser(newUser.Email) {
		util.SendData(w, "User already exists", http.StatusConflict)
		return
	}

	createUser, err := h.svc.Create(domain.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		util.SendData(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createUser, http.StatusCreated)
}
