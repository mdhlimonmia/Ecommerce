package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	id, err := strconv.Atoi(r.PathValue("productId"))
	if err != nil {
		http.Error(w, "Please Give Correct Input", 400)
		return
	}

	fmt.Println("Product id hit: ", id)
	p, ok := database.Update(id, newProduct)
	if !ok {
		http.Error(w, "Product not found", 404)
		return
	}

	util.SendData(w, p, 200)
}
