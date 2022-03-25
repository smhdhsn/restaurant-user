up:
	./script/docker_up.sh $(ENV)
bash:
	docker exec -it bookstore_app bash
build:
	go build -o $(BIN_DIR)/ ./cmd/api
.PHONY: up bash build
