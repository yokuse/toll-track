obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver:
# since we have multiple files, we cannot just build main.go, we build the whole dir
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

proto:
	protoc --proto_path=./types --go-grpc_out=. --go_out=paths=import:. ./types/*.proto

gateway:
	@go build -o bin/gateway ./gateway
	@./bin/gateway

.PHONY: obu aggregator gateway