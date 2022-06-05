package main

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"iteung-api/config"
	"iteung-api/controller"
	"iteung-api/exception"
	"iteung-api/helper"
	"iteung-api/repository"
	"iteung-api/service"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(nice.Recovery(exception.ErrorHandler))

	// Setup Config
	siapDB := config.NewSiapDB()
	validate := validator.New()

	// Setup Repository
	loginRepository := repository.NewLoginRepository()

	// Setup Service
	loginService := service.NewLoginService(loginRepository, siapDB, validate)

	// Setup Controller
	loginController := controller.NewLoginController(loginService)

	// Setup Route
	loginController.Route(r)

	err := r.Run()
	helper.PanicIfError(err)
}
