package service

import (
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func connToElastic() (*elasticsearch.Client, error) {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://64.23.174.193:9200",
		},
		Username: os.Getenv("ELASTIC_USERNAME"), //"elastic",
		Password: os.Getenv("ELASTIC_PASSWORD"), //"Bm4puJHi2C0aR56nhXY5",
	}
	elasticsearchDefaultClient, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, err
	}

	return elasticsearchDefaultClient, nil
}

func seedElastic(elasticClient *elasticsearch.Client) error {
	_, err := elasticClient.Indices.Create("advertisement")
	if err != nil {
		return err
	}
	return nil
}

func NewServiceServer() error {
	elasticClient, err := connToElastic()
	if err != nil {
		return err
	}
	log.Println("Service connected to Elasticsearch")
	log.Println(elasticsearch.Version)
	log.Println(elasticClient.Info())

	err = seedElastic(elasticClient)
	if err != nil {
		return err
	}
	log.Println("Seeded Elasticsearch")

	gRPCListener, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}
	log.Println("Service listening on :8000")

	gRPCServer := grpc.NewServer()
	proto.RegisterServiceServer(gRPCServer, &ServiceServer{elasticClient: elasticClient})

	reflection.Register(gRPCServer)

	return gRPCServer.Serve(gRPCListener)
}
