# syntax=docker/dockerfile:1.1.3-experimental
FROM golang:1.20 as builder
ARG GOPROXY

ENV GOPROXY=${GOPROXY}

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY consent/consent-storage .
COPY shared shared
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM alpine:3.18
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
