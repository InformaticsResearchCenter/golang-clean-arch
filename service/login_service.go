package service

import (
	"github.com/gin-gonic/gin"
	"iteung-api/api"
)

type LoginService interface {
	LoginAuthWithPhoneNumberAndPassword(c *gin.Context, request api.LoginRequest) api.LoginResponse
}
