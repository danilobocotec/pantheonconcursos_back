# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS build

WORKDIR /app

RUN apk add --no-cache build-base git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

FROM alpine:3.19

WORKDIR /app

RUN adduser -D -g '' app

COPY --from=build /app/bin/api /usr/local/bin/api

EXPOSE 8080

ENV GIN_MODE=release

USER app

CMD ["api"]
