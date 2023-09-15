SERVER_IMAGE = wisdom-server
CLIENT_IMAGE = wisdom-client

# Docker compose default action
.PHONY: all
all:
	docker-compose up

# Wisdom Server Targets
.PHONY: build-server
build-server:
	docker build -t $(SERVER_IMAGE) -f cmd/server/Dockerfile .

.PHONY: run-server
run-server: build-server
	docker run --name $(SERVER_IMAGE) -p 8080:8080 $(SERVER_IMAGE)

# Wisdom Client Targets
.PHONY: build-client
build-client:
	docker build -t $(CLIENT_IMAGE) -f cmd/client/Dockerfile .

.PHONY: run-client
run-client: build-client
	docker run -e SERVER_URL=wisdom-server:8080 --link $(SERVER_IMAGE) $(CLIENT_IMAGE)

# Clean-up
.PHONY: clean
clean:
	docker container rm -f $(SERVER_IMAGE) $(CLIENT_IMAGE)
	docker image rm $(SERVER_IMAGE) $(CLIENT_IMAGE)

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all           - Build and run both the server and client using Docker Compose."
	@echo "  build-server  - Build the Wisdom Server Docker image."
	@echo "  run-server    - Run the Wisdom Server Docker container."
	@echo "  build-client  - Build the Wisdom Client Docker image."
	@echo "  run-client    - Run the Wisdom Client Docker container."
	@echo "  clean         - Remove both server and client containers and images."
