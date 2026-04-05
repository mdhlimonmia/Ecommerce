package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid Product ID", 400)
		return
	}
	fmt.Println("productId: ", id)

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}
	http.Error(w, "Product Not Found", 404)
}
