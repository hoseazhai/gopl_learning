package main

import (
	pb "github.com/hoseazhai/gopl_learning/tag_service/proto"
	"github.com/hoseazhai/gopl_learning/tag_service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", "0.0.0.0:"+"8008")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("server.Server err : %v", err)
	}

}
