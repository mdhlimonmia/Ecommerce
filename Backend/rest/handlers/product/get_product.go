package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid Product ID", 400)
		return
	}
	fmt.Println("productId: ", id)

	p, found := database.GetProductById(id)
	if found {
		util.SendData(w, p, 200)
		return
	}

	http.Error(w, "Product Not Found", 404)
}
