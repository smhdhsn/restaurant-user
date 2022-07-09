APP_MODE ?= local

# runs the script which loads the containers of the application.
up:
	@./scripts/docker_up.sh $(APP_MODE)

# deletes application's containers.
purge:
	@docker rm -f restaurant_user_app restaurant_user_db
	@docker volume rm restaurant_user

# accesses the shell of application's container.
shell:
	@docker exec -it restaurant_user_app bash

# builds server's http entry point.
build-server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds all the entry points of the application.
build-all: build-server

# compiles proto files related to user auth.
proto-auth:
	@protoc --go_out=internal/protos/user/auth/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/user/auth/ protos/user/auth/*.proto

# compiles all proto files.
proto-all: proto-auth

.PHONY: up purge shell build-server build-all proto-auth proto-all