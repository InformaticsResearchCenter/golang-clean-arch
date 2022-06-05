package repository

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"iteung-api/model"
)

type LoginRepository interface {
	FindByPhoneNumberAndPassword(c *gin.Context, tx *sql.Tx, phoneNumber string, password string) (model.UserInfo, error)
	FindByPhoneNumber(c *gin.Context, tx *sql.Tx, phoneNumber string) (model.UserInfo, error)
}
