package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w, r)

	// if r.Method != "POST" {
	// 	http.Error(w, "Only allow post method", 400)
	// 	return
	// }

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please Give Correct Input", 400)
		return
	}
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	// w.WriteHeader(201)
	// encoder := json.NewEncoder(w)
	// encoder.Encode(newProduct)
	util.SendData(w, newProduct, 201)
}
