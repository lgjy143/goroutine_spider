package scheduler

import "crawl_zhenai/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request)  {
	go func() {s.workerChan <- request}()
}

func (s *SimpleScheduler) Run()  {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request{
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request)  {

}
