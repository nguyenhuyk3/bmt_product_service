package middlewares

import (
	"bmt_product_service/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	x_user_email = "X-User-Email"
)

type GetFromHeaderMiddleware struct {
}

func NewGetFromHeaderMiddleware() *GetFromHeaderMiddleware {
	return &GetFromHeaderMiddleware{}
}

func (g *GetFromHeaderMiddleware) GetEmailFromHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetHeader(x_user_email)
		if email == "" {
			responses.FailureResponse(c, http.StatusBadRequest, "email is not empty!!")
			c.Abort()
			return
		}

		c.Set(x_user_email, email)
		c.Next()
	}
}
