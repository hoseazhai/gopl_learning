package main

import (
	"context"
	pb "github.com/hoseazhai/gopl_learning/tag_service/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	ctx := context.Background()
	clientConn, err := GetClientConn(ctx, "localhost:8004", nil)
	if err != nil {
		log.Fatalf("err：%v", err)
	}
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTaglist err: %v", err)
	}
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
