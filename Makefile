up:
	./script/docker_up.sh $(APP_MODE)
bash:
	docker exec -it restaurant_user_app bash
build:
	go build -o $(BIN_DIR)/ ./cmd/server
proto_source:
	protoc protos/user/source/*.proto --go_out=plugins=grpc:internal/
.PHONY: up bash build proto_source
