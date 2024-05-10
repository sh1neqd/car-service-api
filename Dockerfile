FROM golang:1.22-alpine AS builder

WORKDIR /usr/src/app

COPY . .

EXPOSE 8000

CMD cd ./app && go run ./cmd/app/main.go

RUN apk --no-cache add bash gettext