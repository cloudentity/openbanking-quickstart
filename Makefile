export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

.PHONY: build
build:
	docker-compose -f docker-compose.yaml build

.PHONY: run
run:
	docker-compose -f docker-compose.yaml up -d

.PHONY: lint
lint:
	golangci-lint run --fix --deadline=5m ./...

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	docker-compose down -v
