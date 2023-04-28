DOCKER_COMPOSE=docker-compose.yml
DOCKER_FILE_SERVER=Dockerfile.server
DOCKER_FILE_CLIENT=Dockerfile.client

.PHONY: build run stop clean

build:
	docker-compose -f $(DOCKER_COMPOSE) build

run:
	docker-compose -f $(DOCKER_COMPOSE) up -d

stop:
	docker-compose -f $(DOCKER_COMPOSE) down

clean:
	docker-compose -f $(DOCKER_COMPOSE) down
	docker image rm pow-server:latest
	docker image rm pow-client:latest