package product

import (
	"ecommerce/product"
	"ecommerce/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc         product.ProductService
}

func NewHandler(middlewares *middleware.Middlewares, svc product.ProductService) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
	}
}
