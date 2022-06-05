package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"iteung-api/api"
	"iteung-api/exception"
	"iteung-api/helper"
	"iteung-api/repository"
)

type LoginServiceImpl struct {
	LoginRepository repository.LoginRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewLoginService(loginRepository repository.LoginRepository, DB *sql.DB, validate *validator.Validate) LoginService {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *LoginServiceImpl) LoginAuthWithPhoneNumberAndPassword(c *gin.Context, request api.LoginRequest) api.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.LoginRepository.FindByPhoneNumber(c, tx, request.PhoneNumber)
	if err != nil {
		panic(exception.NewLoginError(err.Error()))
	}

	userInfo, err := service.LoginRepository.FindByPhoneNumberAndPassword(c, tx, request.PhoneNumber, request.Password)
	if err != nil {
		panic(exception.NewLoginError(err.Error()))
	}

	return api.ToLoginResponse(userInfo)
}
