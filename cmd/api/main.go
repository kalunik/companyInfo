package main

import (
	"github.com/kalunik/companyInfo/internal"
	"log"
)

func main() {
	server := new(internal.HttpServer)
	mux := internal.NewRouter()
	err := server.Run("8080", mux)
	if err != nil {
		log.Fatalf("error occured while running http server: %s\n", err.Error())
	}
}
