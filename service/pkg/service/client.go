package service

import (
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient() (proto.ServiceClient, error) {
	conn, err := grpc.Dial("172.17.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return proto.NewServiceClient(conn), nil
}
