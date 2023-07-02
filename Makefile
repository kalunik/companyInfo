MAIN = cmd/main.go
PROTO_PATH = api
GENERATED_PATH = internal/grpc/proto
CONFIG = config/config_sample.env

docker:
	docker build -t 'company-info' .
	docker run --env-file=$(CONFIG) -p 8080:8080 --rm company-info
#try different ports and delete rm tag after adding graceful stop
local:
	@go run $(MAIN)

googleapi:
	@mkdir -p $(PROTO_PATH)/google/api/
	@echo "Download annotations.proto"
	@curl --request GET -sL \
	     --url 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto'\
	     --output $(PROTO_PATH)/google/api/annotations.proto
	@echo "Download annotations.proto"
	@curl --request GET -sL \
	     --url 'https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto'\
	     --output $(PROTO_PATH)/google/api/http.proto
	@echo "Completed"

proto:
	@protoc \
	--go_out=$(GENERATED_PATH) \
	--go-grpc_out=$(GENERATED_PATH) \
	--grpc-gateway_out=$(GENERATED_PATH) \
	--openapiv2_out=$(GENERATED_PATH) \
	--proto_path=$(PROTO_PATH) company.proto
	@echo "Completed"

.PHONY : docker local proto googleapi