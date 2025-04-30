FROM golang:1.24.2 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.mod go.sum ./

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./proto ./proto
COPY .env ./

RUN protoc --go_out=./proto --go-grpc_out=./proto ./proto/*.proto

RUN go build -o app ./cmd/vote/main.go

EXPOSE 5003

CMD ["./app"]