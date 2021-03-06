ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder
ENV MYSQL_PASSWORD=$(MYSQL_PASSWORD)
ENV MYSQL_USER=$(MYSQL_USER)
ENV MYSQL_HOST=$(MYSQL_HOST)
ENV MYSQL_DBNAME=$(MYSQL_DBNAME)
ENV MYSQL_PORT=$(MYSQL_PORT)

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 3030

ENTRYPOINT ["./app"]