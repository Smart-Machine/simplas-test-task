package main

import (
	"github.com/Smart-Machine/simplas-test-task/service/pkg/service"
	"log"
)

func main() {
	err := service.NewServiceServer()
	if err != nil {
		log.Fatal(err)
	}
}
