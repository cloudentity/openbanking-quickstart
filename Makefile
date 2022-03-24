export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

OB_APPS=configuration developer-tpp financroo-tpp consent-page consent-self-service-portal consent-admin-portal bank
ACP_APPS=acp crdb redis
ACP_ONLY_APPS=acp crdb redis
CDR_ACP_CONFIG_APPS=configuration-cdr

# obuk, obbr, cdr
run-%-local: 
	cp -f .env-local .env
	docker-compose -f docker-compose.acp.local.yaml up -d --no-build ${ACP_APPS}
	./scripts/wait.sh 
	docker-compose -f docker-compose.$*.yaml up --no-build -d 
	./scripts/wait.sh

run-%-saas:
	cp -f .env-saas .env
	docker-compose -f docker-compose.$*.yaml up --no-build -d
	./scripts/wait.sh

.PHONY: build
build:
	docker-compose -f docker-compose.build.yaml build

# obuk, obbr, cdr, saas
run-%-tests-headless: run-tests-verify
	yarn --cwd tests run cypress run -s cypress/integration/$*/*.ts

.PHONY: run-tests
run-tests:
	yarn --cwd tests run cypress open

.PHONY: run-tests-verify 
run-tests-verify: 
	VERIFY_TEST_RUNNER_TIMEOUT_MS=80000 yarn --cwd tests run cypress verify

.PHONY: run-apps
run-apps:
	docker-compose up -d --no-build ${OB_APPS}
	docker-compose -f docker-compose.cdr.yaml up -d ${CDR_APPS}
	./scripts/wait.sh


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

.PHONY: clean
clean: 
	docker-compose -f docker-compose.build.yaml down --remove-orphans


.PHONY: purge
purge:
	docker images -a | grep openbanking-quickstart | awk '{print $3}' | xargs docker rmi -f || true

# enable, disable
%-mfa:
	 if [ $* == "enable" ]; then \
        ./scripts/override_env.sh ENABLE_MFA true; \
    else \
        ./scripts/override_env.sh ENABLE_MFA false; \
    fi


.PHONY: set-version
set-version:
	./scripts/override_env.sh VERSION $(shell ./scripts/version.sh)

.PHONY: set-saas-configuration
set-saas-configuration:
	./scripts/override_env.sh TENANT ${SAAS_TENANT_ID}
	./scripts/override_env.sh CONFIGURATION_CLIENT_ID ${SAAS_CLIENT_ID}
	./scripts/override_env.sh CONFIGURATION_CLIENT_SECRET ${SAAS_CLIENT_SECRET}

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

