//go:build wireinject

package injectors

import (
	"bmt_product_service/internal/implementations/message_broker/consummers"
	"bmt_product_service/internal/injectors/provider"

	"github.com/google/wire"
)

func InitFilmUploadConsummer() (*consummers.FilmUploadConsummer, error) {
	wire.Build(
		provider.ProvideQueries,
		consummers.NewFilmUploadConsummer,
	)

	return &consummers.FilmUploadConsummer{}, nil
}
