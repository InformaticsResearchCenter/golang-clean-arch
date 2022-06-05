package controller

import "github.com/gin-gonic/gin"

type LoginController interface {
	LoginAuth(c *gin.Context)
	Route(r *gin.Engine)
}
