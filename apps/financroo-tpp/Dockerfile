# syntax=docker/dockerfile:1.1.3-experimental
FROM node:14 as js-builder

WORKDIR /app/web/app

COPY apps/financroo-tpp/web/app/package.json package.json
COPY apps/financroo-tpp/web/app/package-lock.json package-lock.json

RUN npm install --no-audit --prefer-offline

COPY apps/financroo-tpp/web/app/tsconfig.json tsconfig.json
COPY apps/financroo-tpp/web/app/public public
COPY apps/financroo-tpp/web/app/src src

RUN npm run build

FROM golang:1.21.4 as go-builder
ARG GOPROXY

ENV GOPROXY=${GOPROXY}

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY generated generated
COPY shared shared
COPY apps/financroo-tpp .
RUN --mount=type=cache,target=/root/.cache/go-build,mode=777 \
    CGO_ENABLED=0 go build -tags static_all -o main .

FROM cloudentity/alpine:edge-01092023
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=go-builder /app/main .
COPY --from=js-builder /app/web/app/build ./web/app/build

CMD ["./main"]
