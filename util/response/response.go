package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Ng(c *gin.Context, message string) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    1,
		Message: message,
		Data:    nil,
	})
}
