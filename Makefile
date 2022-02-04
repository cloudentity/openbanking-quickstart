export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

OB_APPS=developer-tpp financroo-tpp consent-page-uk consent-page-br consent-page-cdr consent-self-service-portal consent-admin-portal bank-uk bank-br
ACP_APPS=acp crdb redis configuration
ACP_ONLY_APPS=acp crdb redis
CDR_ACP_CONFIG_APPS=configuration-cdr
CDR_CONSENT_APPS=consent-page-cdr consent-self-service-portal consent-admin-portal
CDR_APPS=mock-data-recipient mock-register mock-data-holder

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
	docker-compose -f docker-compose.cdr.yaml up -d ${CDR_APPS}
	./scripts/wait.sh

.PHONY: run-cdr-apps-with-acp-local
run-cdr-apps-with-acp-local:
	docker-compose up -d --no-build ${ACP_ONLY_APPS}
	docker-compose up -d --no-build ${CDR_ACP_CONFIG_APPS}
	docker-compose up -d --no-build ${CDR_CONSENT_APPS}
	docker-compose -f docker-compose.cdr.yaml up -d ${CDR_APPS}

.PHONY: run-cdr-apps-with-saas
run-cdr-apps-with-saas:
	docker-compose up -d --no-build --no-deps ${CDR_ACP_CONFIG_APPS}
	docker-compose up -d --no-build --no-deps ${CDR_CONSENT_APPS}
	docker-compose -f docker-compose.cdr.yaml up -d ${CDR_APPS}

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
	docker-compose down -v --remove-orphans

.PHONY: clean-saas
clean-saas: clean
	./scripts/clean_saas.sh

.PHONY: run-tests
run-tests:
	yarn --cwd tests run cypress open

.PHONY: run-tests-headless
run-tests-headless:
	yarn --cwd tests run cypress verify
	yarn --cwd tests run cypress run

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

.PHONY: generate-cdr-clients
generate-cdr-clients: start-runner
	rm -rf ./openbanking/cdr/banking/*
	docker-compose exec runner sh -c \
	"swagger generate client \
		-f api/cdr/cds_banking.yaml \
		-A banking \
		-t ./openbanking/cdr/banking"

.PHONY: obbr
obbr:
	docker-compose -f docker-compose.yaml -f conformance/docker-compose.obb.yaml -f conformance/docker-compose.fapi.yaml ${cmd}

.PHONY: setup_saas_env
setup_saas_env:
	cp -f .env-saas .env

.PHONY: setup_local_env
setup_local_env:
	cp -f .env-local .env
