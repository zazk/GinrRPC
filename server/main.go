package main

import (
	"context"
	"fmt"
	"grpc_tutorial/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	fmt.Println("Serving App on http://localhost:4040")
	if e := srv.Serve(listener); e != nil {
		fmt.Println("Error Serving App")
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	fmt.Println("RPC Add function was invoked")
	result := &proto.Response{}
	result.Result = request.A + request.B

	return result, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	fmt.Println("RPC Multiply function was invoked")
	result := &proto.Response{}
	result.Result = request.A * request.B

	return result, nil
}
