package repository

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"iteung-api/helper"
	"iteung-api/model"
)

type LoginRepositoryImpl struct{}

func NewLoginRepository() LoginRepository {
	return &LoginRepositoryImpl{}
}

func (repository *LoginRepositoryImpl) FindByPhoneNumberAndPassword(c *gin.Context, tx *sql.Tx, phoneNumber string, password string) (model.UserInfo, error) {
	SQL := "select Login, Nama, Handphone, Email from simak_mst_mahasiswa where Handphone = ? and Password = SUBSTR(MD5(MD5(?)), 1, 10)"

	rows, err := tx.QueryContext(c, SQL, phoneNumber)
	helper.PanicIfError(err)

	defer rows.Close()

	userInfo := model.UserInfo{}
	if rows.Next() {
		err := rows.Scan(&userInfo)
		helper.PanicIfError(err)
		return userInfo, nil
	} else {
		SQL := "select Login, Nama, Handphone, Email from simak_mst_dosen where Handphone = ? and Password = SUBSTR(MD5(MD5(?)), 1, 10)"

		rows, err := tx.QueryContext(c, SQL, phoneNumber)
		helper.PanicIfError(err)

		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.PhoneNumber, &userInfo.EmailAddress)
			helper.PanicIfError(err)
			return userInfo, nil
		} else {
			return userInfo, errors.New("Password Salah.")
		}
	}
}

func (repository *LoginRepositoryImpl) FindByPhoneNumber(c *gin.Context, tx *sql.Tx, phoneNumber string) (model.UserInfo, error) {
	SQL := "select Login, Nama, Handphone, Email from simak_mst_mahasiswa where Handphone = ?"

	rows, err := tx.QueryContext(c, SQL, phoneNumber)
	helper.PanicIfError(err)

	defer rows.Close()

	userInfo := model.UserInfo{}
	if rows.Next() {
		err := rows.Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.PhoneNumber, &userInfo.EmailAddress)
		helper.PanicIfError(err)
		return userInfo, nil
	} else {
		SQL := "select Login, Nama, Handphone, Email from simak_mst_dosen where Handphone = ?"

		rows, err := tx.QueryContext(c, SQL, phoneNumber)
		helper.PanicIfError(err)

		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&userInfo)
			helper.PanicIfError(err)
			return userInfo, nil
		} else {
			return userInfo, errors.New("Nomor Handphone Tidak Ditemukan.")
		}
	}
}
