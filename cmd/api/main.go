package main

import (
	"github.com/kalunik/companyInfo/internal"
	"github.com/kalunik/companyInfo/internal/rest"
	"log"
)

func main() {

	go func() {
		server := new(internal.HttpServer)
		mux := rest.NewRouter()
		err := server.Run("8080", mux)
		if err != nil {
			log.Fatalf("error occured while running http server: %s\n", err.Error())
		}
	}()

	server := new(internal.GrpcServer)
	server.Run()

}
