FROM golang:1.17.2 as base
ENV GO111MODULE=on
WORKDIR /code
RUN go get github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1

ARG UID=1000
ARG GID=1000
ARG USERNAME=cloudentity

USER root

RUN groupadd -fg ${GID} ${USERNAME}
RUN useradd -m -u ${UID} -g ${GID} -s /bin/bash ${USERNAME}

COPY go.mod go.mod
COPY go.sum go.sum

RUN --mount=type=ssh,mode=777 go mod download

# golangci-lint
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
  | bash -s -- -b $GOPATH/bin v1.32.0

USER ${USERNAME}
