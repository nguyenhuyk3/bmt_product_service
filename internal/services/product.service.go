package services

import (
	"bmt_product_service/dto/request"
	"context"
)

type IFilm interface {
	AddFilm(ctx context.Context, arg request.AddProductReq) (int, error)
	GetAllFilms(ctx context.Context) (int, interface{}, error)
}
