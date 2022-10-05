ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

# RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY . .
RUN go build -mod vendor -o ./app ./server.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 8080

ENTRYPOINT ["./app"]