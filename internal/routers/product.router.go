package routers

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/global"
	"bmt_product_service/internal/controllers"
	"bmt_product_service/internal/implementations"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(router *gin.RouterGroup) {
	sqlStore := sqlc.NewStore(global.Postgresql)
	productService := implementations.NewProductService(sqlStore)
	authController := controllers.NewProductService(productService)

	productController := router.Group("/film")
	{
		productController.POST("/add", authController.AddFilm)
	}
}
