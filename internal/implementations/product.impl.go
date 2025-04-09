package implementations

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/dto/request"
	"bmt_product_service/internal/services"
	"context"
	"fmt"
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

// GetAllFilms implements services.IFilm.
func (p *productService) GetAllFilms(ctx context.Context) (int, interface{}, error) {
	films, err := p.SqlStore.Queries.GetAllFilms(ctx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if len(films) == 0 {
		return http.StatusNotFound, nil, fmt.Errorf("no movies found")
	}

	return http.StatusOK, films, nil
}
