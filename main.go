package main

import (
	"context"
	"log"
	"net"

	"github.com/zakiyalmaya/grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{
	calculator.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, req *calculator.Request) (*calculator.Response, error) {
	result := req.Num1 + req.Num2
	return &calculator.Response{Result: float32(result)}, nil
}

func (s *server) Subtract(ctx context.Context, req *calculator.Request) (*calculator.Response, error) {
	result := req.Num1 - req.Num2
	return &calculator.Response{Result: float32(result)}, nil
}

func (s *server) Multiple(ctx context.Context, req *calculator.Request) (*calculator.Response, error) {
	result := req.Num1 * req.Num2
	return &calculator.Response{Result: float32(result)}, nil
}

func (s *server) Divide(ctx context.Context, req *calculator.Request) (*calculator.Response, error) {
	if req.Num2 == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Divide by zero")
	}

	result := float32(req.Num1) / float32(req.Num2)
	return &calculator.Response{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{
		calculator.UnimplementedCalculatorServiceServer{},
	})
	log.Println("Server running on port :5000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
