FROM golang:1.15.2 as base
ENV GO111MODULE=on
WORKDIR /code
COPY . /code
RUN go get github.com/go-swagger/go-swagger/cmd/swagger@v0.26.1
