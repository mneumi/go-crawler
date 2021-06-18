package scheduler

import "github.com/mneumi/crawler/types"

type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan types.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r types.Request) {
	go func() {
		s.workerChan <- r
	}()
}
