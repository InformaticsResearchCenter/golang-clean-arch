package exception

import (
	"github.com/gin-gonic/gin"
	"iteung-api/api"
	"iteung-api/helper"
	"net/http"
)

func LoginError(c *gin.Context, error string) {
	apiResponse := api.ResponseAPI{
		Code:    http.StatusUnauthorized,
		Success: false,
		Status:  error,
	}
	helper.WriteToResponseBody(c, http.StatusUnauthorized, apiResponse)
}
