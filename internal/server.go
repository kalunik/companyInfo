package internal

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kalunik/companyInfo/internal/grpc/company"
	gen "github.com/kalunik/companyInfo/internal/grpc/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Grpc *GrpcServer
	Http *GatewayServer
}

func (s *Server) NewServer() *Server {
	return &Server{
		Grpc: &GrpcServer{},
		Http: &GatewayServer{},
	}
}

type GrpcServer struct {
	grpc *grpc.Server
}

func (s *GrpcServer) Run() error {
	grpcServHost := "127.0.0.1"
	grpcServPort := "8080"
	servAddress := net.JoinHostPort(grpcServHost, grpcServPort)
	listen, err := net.Listen("tcp", servAddress)
	if err != nil {
		log.Fatalln("Failed to listen:", err.Error())
	}

	s.grpc = grpc.NewServer()
	gen.RegisterCompanyServer(s.grpc, &company.CompanyGRPC{})

	log.Println("Serving gRPC on", servAddress)
	return s.grpc.Serve(listen)
}

func (s *GrpcServer) Shutdown(ctx context.Context) {
	s.grpc.Stop()
}

type GatewayServer struct {
	gw   *http.Server
	conn *grpc.ClientConn
	mux  *runtime.ServeMux
}

func (s *GatewayServer) ConnectToGRPC() error {
	var err error
	grpcServAddres := net.JoinHostPort("127.0.0.1", "8080")
	s.conn, err = grpc.DialContext(
		context.Background(),
		grpcServAddres,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	log.Println("Connected to gRPC server by address :", grpcServAddres)
	return err
}

func (s *GatewayServer) RegisterHandler() error {
	s.mux = runtime.NewServeMux()
	grpcServAddres := net.JoinHostPort("127.0.0.1", "8080")
	option := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	log.Println("Registring handler on ", grpcServAddres)
	return gen.RegisterCompanyHandlerFromEndpoint(context.Background(), s.mux, grpcServAddres, option)
}

func (s *GatewayServer) Run(port string) error {
	s.gw = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        s.mux,
	}
	log.Println("Serving gRPC-Gateway on", net.JoinHostPort("127.0.0.1", port))
	return s.gw.ListenAndServe()
}

func (s *GatewayServer) Shutdown(ctx context.Context) error {
	return s.gw.Shutdown(ctx)
}
