package main

import (
	"github.com/kalunik/companyInfo/config"
	"github.com/kalunik/companyInfo/internal"
	"log"
)

func main() {

	conf := new(config.Config)
	err := conf.ParseConfig()
	if err != nil {
		log.Fatalln("Parse config from env failed:", err)
	}

	serv := new(internal.Server)
	serv = serv.NewServer()

	go func() {

		err = serv.Http.RegisterHandler()
		if err != nil {
			log.Fatalln("Failed to register the handle :", err)
		}
		err = serv.Http.Run("8081")
		if err != nil {
			log.Fatalln("Failed to run gRPC-Gateway :", err)
		}

	}()

	//err := serv.Http.ConnectToGRPC()
	//if err != nil {
	//	log.Fatalln("Failed to dial server:", err)
	//}
	err = serv.Grpc.Run()
	if err != nil {
		log.Fatalln("Failed to run gRPC server :", err)
	}

	//mux reg list serv

}
