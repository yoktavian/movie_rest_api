package response

import (
	"fmt"
	"movies/entity"

	"github.com/gin-gonic/gin"
)

type GeneralPartnerResponseCode int

const (
	Successfull GeneralPartnerResponseCode = 200
	BadRequest  GeneralPartnerResponseCode = 400
)

var (
	responseMessage = map[GeneralPartnerResponseCode]string{
		Successfull: "Success",
		BadRequest:  "Bad Request",
	}
)

func (rc GeneralPartnerResponseCode) Message() string {
	message, exist := responseMessage[rc]
	if !exist {
		return "Other error"
	}

	return message
}

func PublishCustomError(c *gin.Context, rc GeneralPartnerResponseCode, message string) {
	responData := entity.Response{
		RC:      fmt.Sprintf("%v", rc),
		Message: message,
	}

	c.JSON(int(rc), responData)
}

func Publish(c *gin.Context, rc GeneralPartnerResponseCode, data interface{}) {
	message := rc.Message()

	responData := entity.Response{
		RC:      fmt.Sprintf("%v", rc),
		Message: message,
		Data:    data,
	}

	c.JSON(int(rc), responData)
}
