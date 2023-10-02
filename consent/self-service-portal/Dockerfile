# syntax=docker/dockerfile:1.1.3-experimental
FROM node:14.15.1 as js-builder

WORKDIR /app/web/app

COPY consent/self-service-portal/web/app/package.json package.json
COPY consent/self-service-portal/web/app/package-lock.json package-lock.json

RUN npm install --no-audit --prefer-offline

COPY consent/self-service-portal/web/app/tsconfig.json tsconfig.json
COPY consent/self-service-portal/web/app/public public
COPY consent/self-service-portal/web/app/src src

RUN npm run build

FROM golang:1.20 as go-builder
ARG GOPROXY

ENV GOPROXY=${GOPROXY}

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY generated generated
RUN go mod download

COPY consent/self-service-portal .
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM cloudentity/alpine:edge-01092023
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=go-builder /app/main .
COPY --from=js-builder /app/web/app/build ./web/app/build

CMD ["./main"]
