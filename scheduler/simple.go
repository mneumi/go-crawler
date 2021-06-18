package scheduler

import "github.com/mneumi/crawler/types"

type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) Submit(r types.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan types.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.Request)
}

// 每个worker共用一个channel（Simple），还是每个worker一个channel（Queue）(Scheduler把握)
