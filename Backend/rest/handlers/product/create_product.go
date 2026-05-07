package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w, r)

	// if r.Method != "POST" {
	// 	http.Error(w, "Only allow post method", 400)
	// 	return
	// }
	fmt.Println("CreateProduct handler hit")
	var newProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		util.SendData(w, "Please Give Correct Input", http.StatusBadRequest)
		return
	}

	createProduct, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})
	if err != nil {
		util.SendData(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createProduct, http.StatusCreated)
}
