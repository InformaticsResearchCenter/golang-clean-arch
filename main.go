package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"iteung-api/config"
	"iteung-api/controller"
	"iteung-api/helper"
	"iteung-api/repository"
	"iteung-api/service"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	//r.Use(nice.Recovery(exception.ErrorHandler))
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

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
