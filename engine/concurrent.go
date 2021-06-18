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
	ReadyNotifier
	Submit(types.Request)
	WorkerChan() chan types.Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan types.Request)
}

func (c *ConcurrentEngine) Run(seeds ...types.Request) {
	out := make(chan types.ParseResult)

	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
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

func createWorker(in chan types.Request, out chan types.ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
