package internal

import (
	"context"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Http *HttpServer
	Grpc *grpcServer
}

type grpcServer struct {
	grpc *grpc.Server
}

func (s *grpcServer) Run(port string, mux *chi.Mux) {
	listen, err := net.Listen("tcp", "127.0.0.1:4242")
	if err != nil {
		log.Fatalln("Listen error", err.Error())

	}
	err = s.grpc.Serve(listen)
	if err != nil {
		log.Fatalln("Error while run GRPC server", err.Error())
	}
}

func (s *grpcServer) Shutdown(ctx context.Context) {
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
