#
# See the README.md for instructions how to obtain CLIENT_ID, CLIENT_SECRET and ISSUER_URL.
#
CLIENT_ID	= c7lom1c885j1ctju447g
CLIENT_SECRET	= WdmhAW0lKtqRxPSOnTYIsGA-wuwgNuRC1i7OjO9X1JU
ISSUER_URL	= https://localhost:8443/default/system

format:
	go fmt ./...

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
	PORT=8080 \
	CERT_FILE=../../data/acp_cert.pem \
	KEY_FILE=../../data/acp_key.pem \
	ROOT_CA=../../data/ca.pem \
	FAILURE_URL=https://localhost:8443/default/default/demo?something+went+wrong \
	LOG_LEVEL=debug \
	GIN_MODE=debug \
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
