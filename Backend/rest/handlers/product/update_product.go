package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	id, err := strconv.Atoi(r.PathValue("productId"))
	if err != nil {
		util.SendData(w, "Please Give Correct Input", http.StatusBadRequest)
		return
	}

	fmt.Println("Product id hit: ", id)
	p, err := h.svc.Update(id, domain.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})
	if err != nil {
		util.SendData(w, "Product Not Found", http.StatusNotFound)
		return
	}

	util.SendData(w, p, http.StatusOK)
}
