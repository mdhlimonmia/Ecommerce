package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("productId")

	idInt, err := strconv.Atoi(id) // Convert string to int
	if err != nil {
		http.Error(w, "Invalid Product ID", 400)
		return
	}
	msg := database.DeleteProduct(idInt)
	if msg {
		util.SendData(w, "Product deleted successfully", 200)
		return
	}

	http.Error(w, "Product not found", 404)
}
