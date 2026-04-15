package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w, r)

	// if r.Method != "POST" {
	// 	http.Error(w, "Only allow post method", 400)
	// 	return
	// }
	fmt.Println("CreateProduct handler hit")
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please Give Correct Input", 400)
		return
	}

	id := database.StoreProduct(newProduct)
	newProduct.ID = id

	util.SendData(w, newProduct, 201)
}
