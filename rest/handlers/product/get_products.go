package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

type Pagination struct {
	Data       []*domain.Product `json:"data"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
	TotalItems int               `json:"total_items"`
	TotalPages int               `json:"total_pages"`
}

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	products, err := h.svc.List(int(page), int(limit))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to get products")
		return
	}
	totalItems, err := h.svc.Count()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to get products count")
		return
	}
	totalPages := totalItems / int(limit)
	if totalItems % int(limit) != 0 {
		totalPages++
	}
	paginationData := Pagination{
		Data:       products,
		Page:       int(page),
		Limit:      int(limit),
		TotalItems: 	totalItems,
		TotalPages: totalPages,

	}
	utils.WriteJSON(w, http.StatusOK, paginationData)
}
