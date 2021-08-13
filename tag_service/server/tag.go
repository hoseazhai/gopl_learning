package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hoseazhai/gopl_learning/tag_service/pkg/bapi"
	pb "github.com/hoseazhai/gopl_learning/tag_service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8001")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		fmt.Printf("tttt, : %v \n", err)
		return nil, err
	}
	return &tagList, nil
}
