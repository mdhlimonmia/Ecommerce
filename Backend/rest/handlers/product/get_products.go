package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	productList, err := h.svc.List(limit, page)
	if err != nil {
		util.SendData(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if productList == nil {
		util.SendData(w, "No products found", http.StatusNotFound)
		return
	}

	total, _ := h.svc.TotalProducts() // You need to implement this function to get the total number of products

	util.SendPage(w, productList, limit, page, total)
}
