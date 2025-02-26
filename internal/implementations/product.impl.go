package implementations

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/dto/request"
	"bmt_product_service/internal/services"
	"context"
	"net/http"
)

type productService struct {
	SqlStore *sqlc.SqlStore
}

func NewProductService(sqlStore *sqlc.SqlStore) services.IProduct {
	return &productService{SqlStore: sqlStore}
}

// AddFilm implements services.IProduct.
func (p *productService) AddFilm(ctx context.Context, arg request.AddProductReq) (int, error) {
	err := p.SqlStore.InsertFilmTran(ctx, arg)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
