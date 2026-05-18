package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

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
	var productList []*domain.Product

	wg.Add(1)
	go func() {
		defer wg.Done()
		product, err := h.svc.List(limit, page)
		if err != nil {
			util.SendData(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if product == nil {
			util.SendData(w, "No products found", http.StatusNotFound)
			return
		}
		productList = product
	}()

	var total int

	wg.Add(1)
	go func() {
		defer wg.Done()
		t, err := h.svc.TotalProducts() // You need to implement this function to get the total number of products
		if err != nil {
			util.SendData(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		total = t
	}()

	wg.Wait()
	util.SendPage(w, productList, limit, page, total)
}
