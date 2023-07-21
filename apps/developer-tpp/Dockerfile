# syntax=docker/dockerfile:1.1.3-experimental
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY generated generated
COPY utils utils
COPY apps/developer-tpp .
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/templates templates
COPY --from=builder /app/assets assets
COPY --from=builder /app/main .

CMD ["./main"]
