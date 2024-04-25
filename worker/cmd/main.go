package main

import (
	"context"
	"github.com/Smart-Machine/simplas-test-task/service/pkg/service"
	"github.com/Smart-Machine/simplas-test-task/worker/internal/worker"
	"github.com/Smart-Machine/simplas-test-task/worker/pkg/models/stream"
	"log"
)

const (
	dataFilepath = "./data.json"
	numOfWorkers = 3
)

func main() {

	client, err := service.NewServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	pool := worker.NewPool(numOfWorkers, client)
	errgroup := pool.StartPool(ctx)
	jsonStream := stream.NewJSONStream()
	go func() {
		for data := range jsonStream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
			}
			pool.SendData(data.Advertisement)
		}
	}()

	jsonStream.Start(dataFilepath)
	cancel()
	errgroup.Wait()
}
