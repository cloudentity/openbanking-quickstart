# syntax=docker/dockerfile:1.2
FROM golang:1.21 as builder
ARG GOPROXY

ENV GOPROXY=${GOPROXY}

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY apps/custom-idp .
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM cloudentity/alpine:edge-01092023
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
