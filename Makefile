MAIN = cmd/api/main.go
PROTO_PATH = internal/grpc/proto
GENERATED_PATH = internal/grpc/proto

all:
	@go run $(MAIN)

googleapi:
	mkdir -p $(PROTO_PATH)/google/api/
	curl --request GET -sL \
	     --url 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto'\
	     --output $(PROTO_PATH)/google/api/annotations.proto
	curl --request GET -sL \
	     --url 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto'\
	     --output $(PROTO_PATH)/google/api/http.proto

proto:
	protoc \
	--go_out=$(GENERATED_PATH) \
	--go-grpc_out=$(GENERATED_PATH) \
	--grpc-gateway_out=$(GENERATED_PATH) \
	--openapiv2_out=$(GENERATED_PATH) \
	--proto_path=$(PROTO_PATH) company.proto

.PHONY : all proto googleapi