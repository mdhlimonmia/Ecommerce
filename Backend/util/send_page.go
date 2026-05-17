package util

import "net/http"

type Pagination struct {
	Limit      int `json:"limit"`
	Page       int `json:"page"`
	TotalItems int `json:"total_items"`
	TotalPage  int `json:"total_pages"`
}

type Response struct {
	Data any        `json:"data"`
	Page Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, data any, limit, page, totalItems int) {
	response := Response{
		Data: data,
		Page: Pagination{
			Limit:      limit,
			Page:       page,
			TotalItems: totalItems,
			TotalPage:  totalItems/limit + 1,
		},
	}

	SendData(w, response, http.StatusOK)
}
