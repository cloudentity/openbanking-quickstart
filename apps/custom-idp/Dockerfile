# syntax=docker/dockerfile:1.2
FROM golang:1.17 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY apps/custom-idp .
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
