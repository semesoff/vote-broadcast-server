FROM golang:1.23-bookworm AS builder

# установка рабочей директории
WORKDIR /app

# установка необходимых пакетов для компиляции proto-файлов
RUN apt-get update && apt-get install -y protobuf-compiler

# установка необходимых пакетов для работы с gRPC
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# копирование go.mod и go.sum
COPY go.mod go.sum ./

# загрузка зависимостей
RUN go mod download

# копирование файлов приложения
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./proto ./proto
COPY .env ./

# компиляция proto-файлов
RUN protoc --go_out=./proto --go-grpc_out=./proto ./proto/*.proto

# компиляция go-приложения
RUN go build -o app ./cmd/gateway/main.go

# запуск приложения
CMD ["./app"]