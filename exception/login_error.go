package exception

import (
	"github.com/gin-gonic/gin"
	"iteung-api/api"
	"iteung-api/helper"
	"net/http"
)

type LoginError struct {
	Error string
}

func NewLoginError(c *gin.Context, error string) LoginError {
	apiResponse := api.ResponseAPI{
		Code:    http.StatusUnauthorized,
		Success: false,
		Status:  error,
	}
	helper.WriteToResponseBody(c, http.StatusUnauthorized, apiResponse)
	return LoginError{Error: error}
}
