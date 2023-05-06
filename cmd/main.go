package main

import (
	"github.com/kalunik/companyInfo/internal"
)

func main() {
	//go func() {
	//	server := new(internal.HttpServer)
	//	mux := rest.NewRouter()
	//	err := server.Run("8080", mux)
	//	if err != nil {
	//		log.Fatalf("error occured while running http server: %s\n", err.Error())
	//	}
	//}()

	//server := new(internal.GrpcServer)
	//server.Run()

	//serv := &internal.Server{
	//	Grpc: &internal.GrpcServer{},
	//}
	serv := new(internal.Server)
	serv = serv.NewServer()
	serv.Grpc.Run()
	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	"0.0.0.0:8080",
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	//if err != nil {
	//	log.Fatalln("Failed to dial server:", err)
	//}
	//
	//gwmux := runtime.NewServeMux()
	//// Register Greeter
	//err = gen.RegisterCompanyHandler(context.Background(), gwmux, conn)
	//if err != nil {
	//	log.Fatalln("Failed to register gateway:", err)
	//}
	//
	//gwServer := &http.Server{
	//	Addr:    ":8090",
	//	Handler: gwmux,
	//}
	//
	//log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	//log.Fatalln(gwServer.ListenAndServe())

}
