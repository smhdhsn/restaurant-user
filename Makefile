up:
	./script/docker_up.sh $(APP_MODE)
bash:
	docker exec -it bookstore_app bash
build:
	go build -o $(BIN_DIR)/ ./cmd/api
.PHONY: up bash build
