export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

ARTIFACTORY_USER              ?= jenkins
ARTIFACTORY_GO_PROXY_PASSWORD ?= PUT_ARTIFACTORY_GO_PROXY_PASSWORD_HERE

.EXPORT_ALL_VARIABLES: ;
ACP_LOCAL_APPS      =acp crdb redis jaeger node-env
ALL_DOCKER_COMPOSES =-f docker-compose.obuk.yaml -f docker-compose.obbr.yaml -f docker-compose.cdr.yaml -f docker-compose.generic.yaml -f docker-compose.fdx.yaml -f docker-compose.build.yaml

ifeq ($(CI), true)
	DOCKER_BUILD_PARAMS := --build-arg GOPROXY=https://${ARTIFACTORY_USER}:${ARTIFACTORY_GO_PROXY_PASSWORD}@artifactory.cloudentity.com/artifactory/api/go/go-virtual
endif

.PHONY: build
build:
	cp -f .env-local .env
	docker-compose ${ALL_DOCKER_COMPOSES} build $(DOCKER_BUILD_PARAMS)

# developer-tpp, financroo-tpp, consent-page, bank, configuration, consent-self-service-portal, consent-admin-portal
build-%:
	docker-compose ${ALL_DOCKER_COMPOSES} build $*

# obuk, obbr, cdr, fdx, generic
run-%-local:
	./scripts/additional_configuration.sh $* "local" ${BRANCH_NAME}
	cp -f .env-local .env
	docker-compose -f docker-compose.acp.local.yaml up -d --no-build crdb
	./scripts/wait.sh
	docker-compose -f docker-compose.acp.local.yaml up -d --no-build ${ACP_LOCAL_APPS}
	./scripts/wait.sh
	docker-compose -f docker-compose.$*.yaml up --no-build -d
	./scripts/wait.sh

# obuk, obbr, cdr, fdx
run-%-saas: set_saas_configuration
	./scripts/additional_configuration.sh $* "saas" ${BRANCH_NAME}
	cp -f .env-saas .env
	docker-compose -f docker-compose.$*.yaml up --no-build -d
	./scripts/wait.sh

.PHONY: test
test: start-runner
	docker exec quickstart-runner sh -c "go test -failfast -p=4 -parallel=8 ./..."

.PHONY: build-tests
build-tests:
	yarn --cwd tests install

.PHONY: run-tests
run-tests:
	yarn --cwd tests run cypress open

.PHONY: run-tests-verify
run-tests-verify:
	VERIFY_TEST_RUNNER_TIMEOUT_MS=80000 yarn --cwd tests run cypress verify

# obuk, obbr, cdr, fdx, generic
run-%-tests-headless:
	yarn --cwd tests run cypress run -s cypress/integration/$*/*.ts

# obuk, obbr, cdr, fdx
run-saas-%-tests-headless:
	yarn --cwd tests run cypress run -s cypress/integration/saas/$*/*.ts

.PHONY: clean
clean:
	docker-compose -f docker-compose.build.yaml down --remove-orphans
	docker-compose -f docker-compose.cdr.yaml down --remove-orphans || true
ifeq (${DEBUG},true)
	docker ps -a
	rm -fr mount/cdr/*
endif

clean-saas: set_saas_configuration start-runner
	docker exec quickstart-runner sh -c \
    "go run ./scripts/go/clean_saas.go \
        -tenant=${SAAS_TENANT_ID} \
        -cid=${SAAS_CLEANUP_CLIENT_ID} \
        -csec=${SAAS_CLEANUP_CLIENT_SECRET} \
        -spref=${BRANCH_NAME}"
	make stop-runner
	make clean

.PHONY: purge
purge:
	docker images -a | grep openbanking-quickstart | awk '{print $3}' | xargs docker rmi -f || true

.PHONY: clean-dbs
clean-dbs:
	find mount -name "*.db" | xargs rm -f

.PHONY: restart-acp
restart-acp:
	docker-compose -f docker-compose.acp.local.yaml rm -s -f acp
	docker-compose -f docker-compose.acp.local.yaml up -d --no-build acp

.PHONY: lint
lint: start-runner
	docker exec quickstart-runner sh -c "golangci-lint run --timeout 5m ./..."

# enable, disable
%-mfa:
	./scripts/mfa_configuration.sh $*

# enable, disable
%-tls-financroo:
	./scripts/financroo_tls_configuration.sh $*

.PHONY: set-version
set-version:
	./scripts/override_env.sh VERSION $(shell ./scripts/version.sh)

# br, uk
generate-%-integration-spec: start-runner
	./scripts/generate_bank_spec.sh $*

# obuk, obbr, cdr, fdx
generate-%-clients: start-runner
	./scripts/generate_clients.sh $*

.PHONY: obbr-conformance-config
obbr-conformance-config:
	docker-compose -f docker-compose.acp.local.yaml -f conformance/docker-compose.obb.yaml -f conformance/docker-compose.fapi.yaml ${cmd}

.PHONY: bump_acp
bump_acp:
	./scripts/bump_acp.sh

.PHONY: start-runner
start-runner:
	docker build -t quickstart-runner -f build/runner.dockerfile $(DOCKER_BUILD_PARAMS) .
	docker-compose -f docker-compose.acp.local.yaml up -d runner
	docker ps -a

.PHONY: stop-runner
stop-runner:
	docker rm -f quickstart-runner

.PHONY: set_saas_configuration
set_saas_configuration:
	./scripts/set_saas_configuration.sh

.PHONY: pull-docker-images
pull-docker-images:
	docker-compose ${ALL_DOCKER_COMPOSES} pull

.PHONY: list-docker-images
list-docker-images:
	docker-compose ${ALL_DOCKER_COMPOSES} config | yq -r '.services[].image' | grep cloudentity | grep -iv cdr-mock

.PHONY: retag-docker-images
retag-docker-images:
	make -s list-docker-images | ./scripts/retag_docker_images.sh
