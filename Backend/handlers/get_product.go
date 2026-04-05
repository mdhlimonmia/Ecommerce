package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	// handleCors(w)

	// if r.Method != "GET" {
	// 	http.Error(w, "Only allow get method", 400)
	// 	return
	// }

	// encoder := json.NewEncoder(w)
	// encoder.Encode(productList)
	util.SendData(w, database.ProductList, 200)
}
