//go:build wireinject

package injectors

import (
	"bmt_product_service/internal/injectors/provider"
	"bmt_product_service/internal/message_broker/consummers"

	"github.com/google/wire"
)

func InitFilmUploadConsummer() (*consummers.FilmUploadConsummer, error) {
	wire.Build(
		provider.ProvideQueries,
		consummers.NewFilmUploadConsummer,
	)

	return &consummers.FilmUploadConsummer{}, nil
}
