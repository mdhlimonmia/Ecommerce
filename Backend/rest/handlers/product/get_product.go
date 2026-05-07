package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		util.SendData(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	fmt.Println("productId: ", id)

	p, err := h.productRepo.Get(id)
	if err != nil {
		util.SendData(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if p == nil {
		util.SendData(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, p, http.StatusOK)
}
