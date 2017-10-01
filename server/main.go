package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"

	pb "github.com/robinjmurphy/go-grpc-example/server/proto"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if len(req.Input) < 2 {
		err := grpc.Errorf(
			codes.InvalidArgument,
			"at least two numbers must be specified",
		)
		return &pb.Response{}, err
	}

	var res int32 = 1
	for _, i := range req.Input {
		res = res * i
	}

	return &pb.Response{Result: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
