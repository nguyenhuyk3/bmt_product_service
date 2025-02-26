package controllers

import (
	"bmt_product_service/dto/request"
	"bmt_product_service/internal/responses"
	"bmt_product_service/internal/services"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.IProduct
}

func NewProductService(productService services.IProduct) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc *ProductController) AddFilm(c *gin.Context) {
	var req request.AddProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.FailureResponse(c, http.StatusBadRequest, "request is invalid")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	status, err := pc.ProductService.AddFilm(ctx, req)
	if err != nil {
		responses.FailureResponse(c, status, err.Error())
		return
	}

	responses.SuccessResponse(c, http.StatusOK, "add film perform successfully!!", nil)
}
