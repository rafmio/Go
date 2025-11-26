APP_NAME = primitive-http-server
BINARY = server
PORT = 8080

.PHONY: build run test clean docker-build docker-run docker-push help

help:
	@echo "Available commands:"
	@echo "  build           - Build Go binary"
	@echo "  test            - Run tests"
	@echo "  run             - Run server locally"
	@echo "  clean           - Remove binary and Docker container"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-run      - Run Docker container"
	@echo "  docker-push     - Push Docker image (set TAG first)"

build:
	go build -o $(BINARY) primitive-http-server.go

test:
	go test -v ./...

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run: docker-build
	docker run -d --name $(APP_NAME) -p $(PORT):8080 $(APP_NAME):latest

docker-push:
	@echo "Set TAG=yourrepo/$(APP_NAME):v1.0.0 and run again"
	@echo "docker push $(TAG)"

