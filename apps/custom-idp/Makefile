#
# See the README.md for instructions how to obtain these parameters.
#
# Populate these settings from the Custom IDP OAuth credentials.
CLIENT_ID	= c7o7a1f4dqkttnd4356g
CLIENT_SECRET	= h8udwSBfgfGkwUkHVsrlQjPAKDM9Ng6w995HPh_j1BY
ISSUER_URL	= https://host.docker.internal:8443/default/system
CERT_FILE	= ../../data/bank_cert.pem
KEY_FILE	= ../../data/bank_key.pem
ROOT_CA		= ../../data/ca.pem
PORT		= 8080

# Populate these settings from your Web Server Application.

OIDC_CLIENT_ID		= c7ob5jf4dqkttnd43ilg
OIDC_CLIENT_SECRET	= w7GLTQWNVwVu2D9VeNzZDMj-2tDLbxtaEx690NwcD4A
OIDC_ISSUER_URL		= https://host.docker.internal:8443/default/external-authorizer
OIDC_REDIRECT_URL	= https://host.docker.internal:${PORT}/callback
OIDC_CERT_PATH		= ../../data/acp_cert.pem
OIDC_KEY_PATH		= ../../data/acp_key.pem
OIDC_CA_PATH		= ../../data/ca.pem

fmt:
	go fmt ./...

lint:
	golangci-lint run --disable varnamelen --disable tagliatelle

vet:
	go vet ./...

test:	vet
	go test -test.v -cover -test.run '.' ./...

build:	test
	go build

run:	build
	env \
	CLIENT_ID=${CLIENT_ID} \
	CLIENT_SECRET=${CLIENT_SECRET} \
	ISSUER_URL=${ISSUER_URL} \
	PORT=${PORT} \
	CERT_FILE=${CERT_FILE} \
	KEY_FILE=${KEY_FILE} \
	ROOT_CA=${ROOT_CA} \
	FAILURE_URL=https://localhost:8443/default/default/demo?something+went+wrong \
	LOG_LEVEL=debug \
	GIN_MODE=debug \
	OIDC_CLIENT_ID=${OIDC_CLIENT_ID} \
	OIDC_CLIENT_SECRET=${OIDC_CLIENT_SECRET} \
	OIDC_ISSUER_URL=${OIDC_ISSUER_URL} \
	OIDC_REDIRECT_URL=${OIDC_REDIRECT_URL} \
	OIDC_CERT_PATH=${OIDC_CERT_PATH} \
	OIDC_KEY_PATH=${OIDC_KEY_PATH} \
	OIDC_CA_PATH=${OIDC_CA_PATH} \
	./custom-idp


# Run the app with docker-compose
#
docker-build:
	docker build -f Dockerfile -t custom-idp ../..

up:
	docker-compose up --detach --force-recreate custom-idp

logs:
	docker logs --follow custom-idp

down:
	docker-compose rm --stop --force custom-idp


# Consider all build targets that do not include a dot, to be phony.
phony:
	egrep -o '^[a-z0-9-]+:' Makefile | egrep -o '^[^:]+' | sort | xargs echo '.PHONY:' >> Makefile

.PHONY: build docker-build down format logs phony run test up vet
