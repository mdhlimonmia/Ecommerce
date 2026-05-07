package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("productId")

	idInt, err := strconv.Atoi(id) // Convert string to int
	if err != nil {
		util.SendData(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(idInt)
	if err != nil {
		util.SendData(w, "Product deleted successfully", http.StatusOK)
		return
	}

	util.SendData(w, "Product not found", http.StatusNotFound)
}
