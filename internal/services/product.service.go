package services

import (
	"bmt_product_service/dto/request"
	"context"
)

type IProduct interface {
	AddFilm(ctx context.Context, arg request.AddProductReq) (int, error)
	GetFilmById(ctx context.Context) (interface{}, error)
}
