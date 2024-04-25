package main

import (
	"github.com/Smart-Machine/simplas-test-task/service/pkg/service"
	"log"

	"github.com/Smart-Machine/simplas-test-task/httpProxy/internal/api"
)

func main() {
	serviceClient, err := service.NewServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	r := api.SetupRouter(serviceClient)
	err = r.Run(":8001")
	if err != nil {
		log.Fatal(err)
	}
}
