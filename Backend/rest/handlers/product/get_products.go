package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	p, err := h.svc.List()
	if err != nil {
		util.SendData(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if p == nil {
		util.SendData(w, "No products found", http.StatusNotFound)
		return
	}

	util.SendData(w, p, http.StatusOK)

}
