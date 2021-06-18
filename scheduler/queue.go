package scheduler

import "github.com/mneumi/crawler/types"

type QueueScheduler struct {
	requestChan chan types.Request
	workerChan  chan chan types.Request
}

func (s *QueueScheduler) Submit(r types.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkerReady(w chan types.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) WorkerChan() chan types.Request {
	return make(chan types.Request)
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan types.Request)
	s.requestChan = make(chan types.Request)

	go func() {
		var requestQueue []types.Request
		var workerQueue []chan types.Request

		for {
			var activeRequest types.Request
			var activeWorker chan types.Request // 可能是 nil chan，不会被 select 到

			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeWorker = workerQueue[0]
				activeRequest = requestQueue[0]
			}

			select {
			case r := <-s.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-s.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
