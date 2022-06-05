package controller

import (
	"github.com/gin-gonic/gin"
	"iteung-api/service"
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
	//TODO implement me
	panic("implement me")
}
