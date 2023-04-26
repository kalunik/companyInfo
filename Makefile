
MAIN = cmd/api/main.go
PROTO_PATH = internal/grpc/proto
PROTO_OUT = .

all:
	@go run $(MAIN)

proto:
	protoc -I $(PROTO_PATH) company.proto --go-grpc_out=$(PROTO_OUT)