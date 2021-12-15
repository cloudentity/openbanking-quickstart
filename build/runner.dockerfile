FROM golang:1.16 as base
ENV GO111MODULE=on
WORKDIR /code
RUN go get github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1

# golangci-lint
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
  | bash -s -- -b $GOPATH/bin v1.32.0
