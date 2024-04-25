package worker

import (
	"context"
	"github.com/Smart-Machine/simplas-test-task/service/pkg/proto"
	"github.com/Smart-Machine/simplas-test-task/worker/pkg/models/advertisement"
	"golang.org/x/sync/errgroup"
	"log"
)

type Pool struct {
	workers         []Worker
	roundRobinIndex int
}

func NewPool(numOfWorkers int, consumerClient proto.ServiceClient) *Pool {
	workers := []Worker{}
	for i := 0; i < numOfWorkers; i++ {
		workers = append(workers, NewWorker(consumerClient))
	}
	return &Pool{
		workers:         workers,
		roundRobinIndex: 0,
	}
}

func (p *Pool) StartPool(ctx context.Context) *errgroup.Group {
	group, groupCtx := errgroup.WithContext(ctx)
	for i := 0; i < len(p.workers); i++ {
		group.Go(func() error {
			return p.workers[i].StartLoop(groupCtx)
		})
	}
	return group
}

func (p *Pool) SendData(data advertisement.Advertisement) {
	log.Printf("Sending data for processing to %d\n", p.roundRobinIndex)
	p.workers[p.roundRobinIndex].SendData(data)
	p.incrRRI()
}

func (p *Pool) incrRRI() {
	if p.roundRobinIndex == len(p.workers)-1 {
		p.roundRobinIndex = 0
	} else {
		p.roundRobinIndex = p.roundRobinIndex + 1
	}
}
