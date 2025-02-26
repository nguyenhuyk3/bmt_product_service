package implementations

import (
	"bmt_product_service/dto/request"
	"bmt_product_service/internal/services"
	"context"
	"net/http"
)

type productService struct {
}

func NewProductService() services.IProduct {
	return &productService{}
}

// AddFilm implements services.IProduct.
func (p *productService) AddFilm(ctx context.Context, arg request.AddProductReq) (int, error) {
	return http.StatusOK, nil
}
