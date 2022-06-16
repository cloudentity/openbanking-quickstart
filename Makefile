export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;
ACP_LOCAL_APPS=acp crdb redis

# obuk, obbr, cdr
run-%-local: 
	cp -f .env-local .env
	./scripts/additional_configuration.sh $*
	docker-compose -f docker-compose.acp.local.yaml up -d --no-build ${ACP_LOCAL_APPS}
	./scripts/wait.sh 
	docker-compose -f docker-compose.$*.yaml up --no-build -d 
	./scripts/wait.sh

# obuk, obbr
run-%-saas:
	cp -f .env-saas .env
	./scripts/additional_configuration.sh $*
	docker-compose -f docker-compose.$*.yaml up --no-build -d
	./scripts/wait.sh

.PHONY: build
build:
	cp -f .env-local .env
	docker-compose -f docker-compose.obuk.yaml -f docker-compose.obbr.yaml -f docker-compose.cdr.yaml -f docker-compose.build.yaml build

# obuk, obbr, cdr, fdx
run-%-tests-headless:
	yarn --cwd tests run cypress run -s cypress/integration/$*/*.ts

# obuk, obbr, fdx
run-saas-%-tests-headless:
	yarn --cwd tests run cypress run -s cypress/integration/saas/$*/*.ts

.PHONY: run-tests
run-tests:
	yarn --cwd tests run cypress open

.PHONY: run-tests-verify 
run-tests-verify: 
	VERIFY_TEST_RUNNER_TIMEOUT_MS=80000 yarn --cwd tests run cypress verify

.PHONY: restart-acp
restart-acp:
	docker-compose -f docker-compose.acp.local.yaml rm -s -f acp
	docker-compose -f docker-compose.acp.local.yaml up -d --no-build acp

.PHONY: lint
lint: start-runner
	docker exec quickstart-runner sh -c "golangci-lint run --fix --deadline=5m ./..."

.PHONY: clean
clean: 
	docker-compose -f docker-compose.build.yaml down --remove-orphans
ifeq (${DEBUG},true)
	docker ps -a
	rm -fr mount/cdr/*
endif

clean-saas: start-runner
	docker exec quickstart-runner sh -c \
    "go run ./scripts/go/clean_saas.go \
        -tenant=${SAAS_TENANT_ID} \
        -cid=${SAAS_CLEANUP_CLIENT_ID} \
        -csec=${SAAS_CLEANUP_CLIENT_SECRET}"
	make stop-runner
	make clean

.PHONY: purge
purge:
	docker images -a | grep openbanking-quickstart | awk '{print $3}' | xargs docker rmi -f || true

.PHONY: enable-mfa
enable-mfa:
	./scripts/override_env.sh ENABLE_MFA true

.PHONY: disable-mfa
disable-mfa:
	./scripts/override_env.sh ENABLE_MFA false

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
	docker-compose -f docker-compose.acp.local.yaml up -d runner
	docker ps -a

.PHONY: stop-runner
stop-runner:
	docker rm -f quickstart-runner

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
	docker-compose -f docker-compose.acp.local.yaml exec runner sh -c \
	"swagger generate client \
	    --include-tag=obuk
		-f api/obbr/accounts.yaml \
		-A accounts  \
		-t ./openbanking/obbr/accounts"

.PHONY: generate-cdr-clients
generate-cdr-clients: start-runner
	rm -rf ./openbanking/cdr/banking/*
	docker-compose -f docker-compose.acp.local.yaml exec runner sh -c \
	"swagger generate client \
		-f api/cdr/cds_banking.yaml \
		-A banking \
		-t ./openbanking/cdr/banking"

.PHONY: obbr
obbr:
	docker-compose -f docker-compose.acp.local.yaml -f conformance/docker-compose.obb.yaml -f conformance/docker-compose.fapi.yaml ${cmd}

# enable, disable
%-tls-financroo:
	./scripts/financroo_tls_configuration.sh $*
	
.PHONY: bump_acp
bump_acp:
	./scripts/bump_acp.sh
