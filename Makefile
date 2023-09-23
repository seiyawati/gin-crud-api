DOCKER_COMPOSE = docker-compose

API_SERVICE = api
DB_SERVICE = db

build:
	$(DOCKER_COMPOSE) build

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

restart:
	$(DOCKER_COMPOSE) restart

logs:
	$(DOCKER_COMPOSE) logs

shell:
	$(DOCKER_COMPOSE) exec $(API_SERVICE) sh

db-shell:
	$(DOCKER_COMPOSE) exec $(DB_SERVICE) psql -U postgres

test:
	$(DOCKER_COMPOSE) exec $(API_SERVICE) go test ./...

format:
	$(DOCKER_COMPOSE) exec $(API_SERVICE) go fmt ./...

lint:
	$(DOCKER_COMPOSE) exec $(API_SERVICE) staticcheck ./...

swagger-editor:
	open http://localhost:8082

swagger-ui:
	open http://localhost:8081

.PHONY: build up down restart logs shell db-shell test format lint swagger-editor swagger-ui
