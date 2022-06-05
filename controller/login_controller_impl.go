package controller

import (
	"github.com/gin-gonic/gin"
	"iteung-api/api"
	"iteung-api/helper"
	"iteung-api/service"
	"net/http"
)

type LoginConrollerImpl struct {
	LoginService service.LoginService
}

func NewLoginController(loginService service.LoginService) LoginController {
	return &LoginConrollerImpl{
		LoginService: loginService,
	}
}

func (controller *LoginConrollerImpl) Route(r *gin.Engine) {
	r.POST("/api/v1/login", controller.LoginAuth)
}

func (contoller *LoginConrollerImpl) LoginAuth(c *gin.Context) {
	loginRequest := api.LoginRequest{}
	helper.ReadFromRequestBody(c, &loginRequest)

	loginResponse := contoller.LoginService.LoginAuthWithPhoneNumberAndPassword(c, loginRequest)
	apiResponse := api.ResponseAPI{
		Code:    http.StatusOK,
		Success: true,
		Status:  "Berhasil Login.",
		Data:    loginResponse,
	}

	helper.WriteToResponseBody(c, http.StatusOK, apiResponse)
}
