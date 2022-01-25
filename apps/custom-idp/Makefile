#
# See the README.md for instructions how to obtain these parameters.
#
# Populate these settings from the Custom IDP OAuth credentials.
CLIENT_ID	= c7neprc3tlitkp4ruaog
CLIENT_SECRET	= jIgg1LtVGBHMp_RxYdi_PgCIqG_90cxT0_NbR8luTXY
ISSUER_URL	= https://localhost:8443/default/system
CERT_FILE	= ../../data/bank_cert.pem
KEY_FILE	= ../../data/bank_key.pem
ROOT_CA		= ../../data/ca.pem
PORT		= 8080

# Populate these settings from your Web Server Application.

OIDC_CLIENT_ID		= c7niscs3tlitkp4rubsg
OIDC_CLIENT_SECRET	= vP-wqzcIowQg1rQrSY_FN_XgumW9MRO6-MqMk6_KFcQ
OIDC_ISSUER_URL		= https://host.docker.internal:8443/default/default
OIDC_REDIRECT_URL	= https://localhost:${PORT}/callback
OIDC_CERT_PATH		= ../../data/acp_cert.pem
OIDC_KEY_PATH		= ../../data/acp_key.pem
OIDC_CA_PATH		= ../../data/ca.pem


fmt:
	go fmt ./...

lint:
	golangci-lint run --disable varnamelen

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

#
# Run the app with docker-compose
#
docker-build:
	docker build -f Dockerfile -t custom-idp ../..

up:
	docker-compose up --detach custom-idp

logs:
	docker logs --follow custom-idp

down:
	docker-compose rm --stop --force custom-idp


# Consider all build targets that do not include a dot, to be phony.
phony:
	egrep -o '^[a-z0-9-]+:' Makefile | egrep -o '^[^:]+' | sort | xargs echo '.PHONY:' >> Makefile

.PHONY: build docker-build down format logs phony run test up vet
