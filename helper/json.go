package helper

import (
	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(c *gin.Context, result interface{}) {
	err := c.BindJSON(result)
	PanicIfError(err)
}

func WriteToResponseBody(c *gin.Context, httpCode int, result interface{}) {
	c.JSON(
		httpCode,
		result,
	)
}
