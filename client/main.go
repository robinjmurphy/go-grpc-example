package main

import (
	"context"
	"log"
	"os"
	"strconv"

	pb "github.com/robinjmurphy/go-grpc-example/server/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func argsToInput(args []string) ([]int32, error) {
	var ret []int32
	for _, a := range args {
		i, err := strconv.Atoi(a)
		if err != nil {
			return ret, err
		}
		ret = append(ret, int32(i))
	}
	return ret, nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	input, err := argsToInput(os.Args[1:])
	if err != nil {
		log.Fatalf("failed to convert args: %v", err)
	}

	r, err := c.Multiply(context.Background(), &pb.Request{
		Input: input,
	})
	if err != nil {
		log.Fatalf("failed to multiply: %v", err)
	}

	log.Printf("result: %v", r.Result)
}
