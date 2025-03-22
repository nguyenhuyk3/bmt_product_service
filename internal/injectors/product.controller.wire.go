//go:build wireinject

package injectors

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/internal/controllers"
	"bmt_product_service/internal/implementations"
	"bmt_product_service/internal/injectors/provider"

	"github.com/google/wire"
)

func InitProductController() (*controllers.ProductController, error) {
	wire.Build(
		provider.ProvidePgxPool,
		sqlc.NewStore,
		implementations.NewProductService,
		controllers.NewProductController,
	)

	return &controllers.ProductController{}, nil
}
