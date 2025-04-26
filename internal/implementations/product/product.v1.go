package product

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/dto/request"
	"bmt_product_service/global"
	"bmt_product_service/internal/services"
	"context"
	"fmt"
	"net/http"
)

type productService struct {
	SqlStore    sqlc.IStore
	RedisClient services.IRedis
}

func NewProductService(
	sqlStore sqlc.IStore,
	redisClient services.IRedis) services.IFilm {
	return &productService{
		SqlStore:    sqlStore,
		RedisClient: redisClient,
	}
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
	var films []sqlc.GetAllFilmsRow

	err := p.RedisClient.Get(global.GET_ALL_FILMS_WITH_ADMIN_ROLE, &films)
	if err != nil {
		if err.Error() == fmt.Sprintf("key %s does not exist", global.GET_ALL_FILMS_WITH_ADMIN_ROLE) {
			films, err = p.SqlStore.GetAllFilms(ctx)
			if err != nil {
				return http.StatusInternalServerError, nil, err
			}

			savingErr := p.RedisClient.Save(global.GET_ALL_FILMS_WITH_ADMIN_ROLE, &films, 60*24*10)
			if savingErr != nil {
				return http.StatusOK, nil, fmt.Errorf("warning: failed to save to Redis: %v", savingErr)
			}

			return http.StatusOK, films, nil
		}

		return http.StatusInternalServerError, nil, fmt.Errorf("getting value occur an error: %v", err)
	}

	return http.StatusOK, films, nil
}
