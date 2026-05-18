package product

func (h *Handler) TotalProducts() (int, error) {
	return h.svc.TotalProducts()
}
