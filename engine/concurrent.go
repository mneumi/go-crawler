package engine

import (
	"log"

	"github.com/mneumi/crawler/model"
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
		if isDuplicate(r.Url) { // seeds 也进行去重的原因是把seed也放入参考表中
			log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		c.Scheduler.Submit(r)
	}

	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("got item #%d: %v", profileCount, item)
				profileCount++
			}
		}

		// 去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				// log.Printf("Duplicate request: %s", request.Url)
				continue
			}

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

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
