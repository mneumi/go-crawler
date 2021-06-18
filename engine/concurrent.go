package engine

import (
	"log"

	"github.com/mneumi/crawler/types"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(types.Request)
	ConfigureMasterWorkerChan(chan types.Request)
}

func (c *ConcurrentEngine) Run(seeds ...types.Request) {
	in := make(chan types.Request)
	out := make(chan types.ParseResult)

	c.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("got item: %v", item)
		}

		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan types.Request, out chan types.ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
