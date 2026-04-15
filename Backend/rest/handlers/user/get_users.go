package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := database.UserList()
	util.SendData(w, userList, http.StatusCreated)
}
