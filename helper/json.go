package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"iteung-api/exception"
)

func ReadFromRequestBody(c *gin.Context, result interface{}) {
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(result)
	exception.PanicIfError(err)
}

func WriteToResponseBody(c *gin.Context, httpCode int, result interface{}) {
	c.JSON(
		httpCode,
		result,
	)
}
