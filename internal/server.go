package internal

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/kalunik/companyInfo/internal/grpc/company"
	gen "github.com/kalunik/companyInfo/internal/grpc/proto/generated"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Grpc *GrpcServer
	Http *HttpServer
}

type GrpcServer struct {
	grpc *grpc.Server
}

func (s *Server) NewServer() *Server {
	return &Server{
		Grpc: &GrpcServer{},
		Http: &HttpServer{},
	}
}

func (s *GrpcServer) Run() {
	grpcServHost := "127.0.0.1"
	grpcServPort := "4242"
	servAddress := net.JoinHostPort(grpcServHost, grpcServPort)
	listen, err := net.Listen("tcp", servAddress)
	if err != nil {
		log.Fatalln("Failed to listen:", err.Error())
	}

	s.grpc = grpc.NewServer()
	gen.RegisterCompanyServer(s.grpc, &company.CompanyGRPC{})

	log.Println("Serving gRPC on", servAddress)
	err = s.grpc.Serve(listen)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *GrpcServer) Shutdown(ctx context.Context) {
	s.grpc.Stop()
}

type HttpServer struct {
	Http *http.Server
}

func (s *HttpServer) Run(port string, mux *chi.Mux) error {
	s.Http = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        mux,
	}
	return s.Http.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.Http.Shutdown(ctx)
}
