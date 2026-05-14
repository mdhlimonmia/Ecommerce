package user

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	userList, err := h.svc.UserList()
	if err != nil {
		util.SendData(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	util.SendData(w, userList, http.StatusCreated)
}
