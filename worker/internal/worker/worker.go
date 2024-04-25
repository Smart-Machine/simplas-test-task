package worker

import (
	"context"
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"github.com/Smart-Machine/simplas-test-task/worker/pkg/models/advertisement"
	"log"
)

type Worker struct {
	data           chan advertisement.Advertisement
	consumerClient proto.ServiceClient
}

func NewWorker(consumerClientConn proto.ServiceClient) Worker {
	return Worker{
		data:           make(chan advertisement.Advertisement),
		consumerClient: consumerClientConn,
	}
}

func (w Worker) StartLoop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			close(w.data)
			return nil
		case d := <-w.data:
			if err := w.processData(ctx, d); err != nil {
				return err
			}
		}
	}
}

func (w Worker) SendData(data advertisement.Advertisement) {
	w.data <- data
}

func (w Worker) processData(ctx context.Context, data advertisement.Advertisement) error {
	res, err := w.consumerClient.Create(ctx, &proto.APICreateRequest{
		Id:         data.ID,
		Categories: data.Categories,
		Title:      data.Title,
		Type:       data.Type,
		Posted:     data.Posted,
	})
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}
