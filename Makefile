export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

OB_APPS=developer-tpp financroo-tpp consent-page consent-self-service-portal consent-admin-portal bank
ACP_APPS=acp crdb hazelcast configuration

.PHONY: build
build:
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml build

.PHONY: run-dev
run-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml up -d
	./scripts/wait.sh

.PHONY: run-acp
run-acp-apps:
	docker-compose up -d --no-build ${ACP_APPS}
	./scripts/wait.sh

.PHONY: stop-acp-apps
stop-acp-apps:
	docker-compose rm -s -f ${ACP_APPS}

.PHONY: run-apps
run-apps:
	docker-compose up -d --no-build ${OB_APPS}

.PHONY: run
run:
	make run-acp-apps run-apps

.PHONY: restart-acp
restart-acp:
	docker-compose rm -s -f acp
	docker-compose up -d --no-build acp

.PHONY: lint
lint:
	golangci-lint run --fix --deadline=5m ./...

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	docker-compose down -v

.PHONY: run-tests
run-tests:
	yarn --cwd tests run cypress open

.PHONY: enable-mfa
enable-mfa:
	./scripts/override_env.sh ENABLE_MFA true

.PHONY: disable-mfa
disable-mfa:
	./scripts/override_env.sh ENABLE_MFA false
