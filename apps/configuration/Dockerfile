# syntax=docker/dockerfile:1.1.3-experimental
FROM golang:1.16 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY data/imports imports

RUN go mod download

COPY apps/configuration .
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/imports imports

CMD ["./main"]
