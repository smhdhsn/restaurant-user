up:
	./script/docker_up.sh $(APP_MODE)
bash:
	docker exec -it restaurant_app bash
build:
	go build -o $(BIN_DIR)/ ./cmd/api
.PHONY: up bash build
