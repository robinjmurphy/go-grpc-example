# go-grpc-example

> A simple example of using [gRPC](https://grpc.io/) in Go

## Installation

```
brew install protoc
go get ./...
```

## Usage

Generate the Go protobuf definition:

```bash
make
```

Start the server:

```bash
go run server/main.go
```

Run the client:

```bash
go run client/main.go 2 4
# 2017/10/01 17:57:40 result: 8
```
