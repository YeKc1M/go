package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcdemo/pb/template_engine"
	"log"
	"net"
)

const (
	PORT = ":9192"
)

type server struct {}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Println("/echo: ", in.Msg)
	return &pb.StringMessage{Msg: "Hello " + in.Msg}, nil
}

func (s *server) Hello(ctx context.Context, in *pb.BasicMessage) (*pb.StringMessage, error) {
	log.Println("/hello")
	return &pb.StringMessage{Msg: "hello world!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestEngineServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}
