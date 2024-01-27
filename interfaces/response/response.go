package response

import (
	"movies/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	responData := entity.Response{
		Status:  "200",
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, responData)
}

func Error(c *gin.Context, status int, message string) {
	responData := entity.Response{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responData)
}
