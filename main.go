package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"iteung-api/helper"
	"net/http"
	"os"
)

func WeOnEnv(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "we are on %s env", os.Getenv("APP_ENV"))
}

func main() {
	router := httprouter.New()

	router.GET("/api/v1/", WeOnEnv)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
