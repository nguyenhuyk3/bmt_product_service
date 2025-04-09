package controllers

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/dto/request"
	"bmt_product_service/global"
	"bmt_product_service/internal/responses"
	"bmt_product_service/internal/services"
	"bmt_product_service/utils/redis"
	"context"
	"fmt"
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

	var result []sqlc.GetAllFilmsRow
	err := redis.Get(global.GET_ALL_FILMS_WITH_ADMIN_ROLE, &result)
	if err != nil {
		if err.Error() == fmt.Sprintf("key %s does not exist", global.GET_ALL_FILMS_WITH_ADMIN_ROLE) {
			status, data, err := pc.ProductService.GetAllFilms(ctx)
			if err != nil {
				responses.FailureResponse(c, status, err.Error())
				return
			}

			saveErr := redis.Save(global.GET_ALL_FILMS_WITH_ADMIN_ROLE, data, 60*24*10)
			if saveErr != nil {
				fmt.Printf("warning: failed to save to Redis: %v\n", saveErr)
			}

			responses.SuccessResponse(c, status, "get all films from DB successfully", data)
			return
		}

		responses.FailureResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, http.StatusOK, "get all films from Redis cache successfully", result)
}
