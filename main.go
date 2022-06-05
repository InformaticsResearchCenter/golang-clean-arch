package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iteung-api/exception"
	"net/http"
	"os"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func WebOnEnv(c *gin.Context) {
	messageResponse := MessageResponse{}
	messageResponse.Message = fmt.Sprintf("we are on %s env", os.Getenv("APP_ENV"))
	c.JSON(http.StatusOK, messageResponse)
}

func main() {
	r := gin.Default()

	r.GET("/api/v1", WebOnEnv)

	err := r.Run()
	exception.PanicIfError(err)
}
