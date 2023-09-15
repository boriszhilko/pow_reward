.PHONY: build-server run-server build-client run-client all clean help

build-server:
	@docker build -t wisdom-server -f cmd/server/Dockerfile .

run-server:
	-docker rm -f wisdom-server
	@docker run --name wisdom-server -p 8080:8080 wisdom-server

build-client:
	@docker build -t wisdom-client -f cmd/client/Dockerfile .

run-client:
	-docker rm -f wisdom-client
	@docker run --link wisdom-server --name wisdom-client -e SERVER_URL=wisdom-server:8080 wisdom-client

all:
	@docker-compose up

clean:
	@docker container prune -f
	@docker image prune -f -a --filter label=stage=builder

help:
	@echo "Available commands:"
	@echo "  make build-server    - Build the server Docker image."
	@echo "  make run-server      - Run the server Docker container."
	@echo "  make build-client    - Build the client Docker image."
	@echo "  make run-client      - Run the client Docker container."
	@echo "  make all             - Use Docker Compose to set up and run both the server and client."
	@echo "  make clean           - Remove all unused Docker containers and builder images."
