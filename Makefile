export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

OB_APPS=developer-tpp financroo-tpp consent-page-uk consent-page-br consent-self-service-portal consent-admin-portal bank-uk bank-br
ACP_APPS=acp crdb hazelcast configuration

.PHONY: build
build:
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml build

.PHONY: run-dev
run-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml up -d
	./scripts/wait.sh

.PHONY: run-acp-apps
run-acp-apps: setup_local_env
	docker-compose up -d --no-build ${ACP_APPS}
	./scripts/wait.sh

.PHONY: stop-acp-apps
stop-acp-apps:
	docker-compose rm -s -f ${ACP_APPS}

.PHONY: run-apps
run-apps:
	docker-compose up -d --no-build ${OB_APPS}

.PHONY: run-apps-with-saas
run-apps-with-saas: setup_saas_env
	docker-compose up -d --no-build configuration
	docker-compose up -d --no-build ${OB_APPS}

.PHONY: run
run:
	make run-acp-apps run-apps

.PHONY: restart-acp
restart-acp:
	docker-compose rm -s -f acp
	docker-compose up -d --no-build acp

.PHONY: lint
lint: start-runner
	docker-compose exec runner sh -c "golangci-lint run --fix --deadline=5m ./..."

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	docker-compose down -v

.PHONY: clean-saas
clean-saas: clean
	./scripts/clean_saas.sh

.PHONY: run-tests
run-tests:
	yarn --cwd tests run cypress open

.PHONY: enable-mfa
enable-mfa:
	./scripts/override_env.sh ENABLE_MFA true

.PHONY: disable-mfa
disable-mfa:
	./scripts/override_env.sh ENABLE_MFA false

enable-spec-obuk:
	./scripts/override_env.sh SPEC obuk
	./scripts/override_env.sh OPENBANKING_SERVER_ID openbanking
	./scripts/override_env.sh DEVELOPER_CLIENT_ID bugkgm23g9kregtu051g
	./scripts/override_env.sh CONSENT_PAGE_CLIENT_ID bv0ocudfotn6edhsiu7g
	./scripts/override_env.sh BANK_CLIENT_ID bukj5p6k7qdmm5ppbi4g
	./scripts/override_env.sh BANK_URL http://bank-uk:8070

enable-spec-obbr:
	./scripts/override_env.sh SPEC obbr 
	./scripts/override_env.sh OPENBANKING_SERVER_ID openbanking_brasil
	./scripts/override_env.sh DEVELOPER_CLIENT_ID bukj5p6k7qdmm5other1
	./scripts/override_env.sh BANK_CLIENT_ID bukj5p6k7qdmm5pother2
	./scripts/override_env.sh CONSENT_PAGE_CLIENT_ID bukj5p6k7qdMIIDfjCCAmagAwImm5ppxxxx
	./scripts/override_env.sh BANK_URL http://bank-br:8070


.PHONY: set-version
set-version:
	./scripts/override_env.sh VERSION $(shell ./scripts/version.sh)

.PHONY: start-runner
start-runner:
	docker build -t quickstart-runner -f build/runner.dockerfile .
	docker-compose up -d runner

.PHONY: generate-openbanking-integration-specs 
generate-openbanking-integration-specs: generate-obuk-integration-spec 

.PHONY: generate-obuk-integration-spec
generate-obuk-integration-spec: start-runner
	./scripts/generate_bank_spec.sh uk

.PHONY: generate-obbr-integration-spec
generate-obbr-integration-spec: start-runner
	./scripts/generate_bank_spec.sh br

.PHONY: generate-integration-specs
generate-integration-specs: generate-obuk-integration-spec generate-obbr-integration-spec

.PHONY: generate-obbr-clients
generate-obbr-clients: start-runner
	rm -rf ./openbanking/obbr/accounts/*
	docker-compose exec runner sh -c \
	"swagger generate client \
	    --include-tag=obuk
		-f api/obbr/accounts.yaml \
		-A accounts  \
		-t ./openbanking/obbr/accounts"

.PHONY: obbr
obbr:
	docker-compose -f docker-compose.yaml -f conformance/docker-compose.obb.yaml -f conformance/docker-compose.fapi.yaml ${cmd}

.PHONY: setup_saas_env
setup_saas_env:
	cp -f .env-saas .env

.PHONY: setup_local_env
setup_local_env:
	cp -f .env-local .env