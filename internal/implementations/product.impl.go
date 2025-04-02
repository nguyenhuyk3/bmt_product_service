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

func NewProductService(sqlStore *sqlc.SqlStore) services.IFilm {
	return &productService{SqlStore: sqlStore}
}

// AddFilm implements services.IFilm.
func (p *productService) AddFilm(ctx context.Context, arg request.AddProductReq) (int, error) {
	err := p.SqlStore.InsertFilmTran(ctx, arg)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// GetFilmById implements services.IFilm.
func (p *productService) GetFilmById(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"name":     "John",
		"age":      30,
		"isActive": true,
		"scores":   []int{85, 90, 78},
		"address": map[string]string{
			"city":    "Hanoi",
			"country": "Vietnam",
		},
	}, nil
}
