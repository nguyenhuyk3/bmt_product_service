package controllers

import (
	"bmt_product_service/dto/request"
	"bmt_product_service/global"
	"bmt_product_service/internal/responses"
	"bmt_product_service/internal/services"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.IFilm
}

func NewProductController(productService services.IFilm) *ProductController {
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

	req.FilmChanges.ChangedBy = c.GetString(global.X_USER_EMAIL)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	status, err := pc.ProductService.AddFilm(ctx, req)
	if err != nil {
		responses.FailureResponse(c, status, err.Error())
		return
	}

	responses.SuccessResponse(c, status, "add film perform successfully", nil)
}

func (pc *ProductController) GetAllFilms(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	status, data, err := pc.ProductService.GetAllFilms(ctx)
	if err != nil {
		responses.FailureResponse(c, status, err.Error())
		return
	}

	responses.SuccessResponse(c, status, "get all films perform successfully", data)
}
