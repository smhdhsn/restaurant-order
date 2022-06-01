.PHONY: up purge shell build_server build_all proto_inventory proto_all

# runs the script which loads the containers of the application.
up:
	@./script/docker_up.sh $(APP_MODE)

# deletes application's containers.
purge:
	@docker rm -f restaurant_order_app restaurant_order_db
	@docker volume rm restaurant_order

# accesses the shell of application's container.
shell:
	@docker exec -it restaurant_order_app bash

# builds server's http entry point.
build_server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds all the entry points of the application.
build_all: build_server

# compiles proto files related to edible inventory.
proto_inventory:
	@protoc --go_out=internal/protos/edible/inventory/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/inventory/ protos/edible/inventory/*.proto

# compiles all proto files.
proto_all: proto_inventory
