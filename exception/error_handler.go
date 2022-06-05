package exception

import (
	"github.com/gin-gonic/gin"
	"iteung-api/api"
	"iteung-api/helper"
	"net/http"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	if notFoundError(c, err) {
		return
	}

	if loginError(c, err) {
		return
	}

	internalServerError(c, err)
}

func notFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		apiResponse := api.ResponseAPI{
			Code:    http.StatusNotFound,
			Success: false,
			Status:  exception.Error,
		}
		helper.WriteToResponseBody(c, http.StatusNotFound, apiResponse)
		return true
	} else {
		return false
	}
}

func loginError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(LoginError)
	if ok {
		apiResponse := api.ResponseAPI{
			Code:    http.StatusUnauthorized,
			Success: false,
			Status:  exception.Error,
		}
		helper.WriteToResponseBody(c, http.StatusUnauthorized, apiResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(c *gin.Context, err interface{}) {
	apiResponse := api.ResponseAPI{
		Code:    http.StatusInternalServerError,
		Success: false,
		Status:  "INTERNAL SERVER ERROR",
		Data:    err,
	}
	helper.WriteToResponseBody(c, http.StatusInternalServerError, apiResponse)
}