generate:
	protoc -I server/proto server/proto/calculator.proto --go_out=plugins=grpc:server/proto

.PHONY:
	generate
