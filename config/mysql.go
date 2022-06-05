package config

import (
	"database/sql"
	"fmt"
	"iteung-api/helper"
	"os"
	"strconv"
	"time"
)

func NewSiapDB() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME_SIAP"),
		os.Getenv("DB_PASSWORD_SIAP"),
		os.Getenv("DB_HOST_SIAP"),
		os.Getenv("DB_PORT_SIAP"),
		os.Getenv("DB_NAME_SIAP"),
	)

	db, err := sql.Open("mysql", dataSourceName)
	helper.PanicIfError(err)

	maxIdleConnection, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	maxOpenConnection, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONN"))
	connMaxLifetime, _ := strconv.Atoi(os.Getenv("CONN_MAX_LIFETIME"))
	connMaxIdletime, _ := strconv.Atoi(os.Getenv("CONN_MAX_IDLETIME"))

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdletime) * time.Minute)

	return db

}

func NewITeungDB() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME_ITEUNG"),
		os.Getenv("DB_PASSWORD_ITEUNG"),
		os.Getenv("DB_HOST_ITEUNG"),
		os.Getenv("DB_PORT_ITEUNG"),
		os.Getenv("DB_NAME_ITEUNG"),
	)

	db, err := sql.Open("mysql", dataSourceName)
	helper.PanicIfError(err)

	maxIdleConnection, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	maxOpenConnection, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONN"))
	connMaxLifetime, _ := strconv.Atoi(os.Getenv("CONN_MAX_LIFETIME"))
	connMaxIdletime, _ := strconv.Atoi(os.Getenv("CONN_MAX_IDLETIME"))

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdletime) * time.Minute)

	return db

}
