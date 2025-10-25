package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "studysync-matchmaking-engine/internal/testpb"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})

	// THIS IS THE KEY LINE
	reflection.Register(s)

	log.Println("listening on :50051")
	s.Serve(lis)
}
