package scheduler

import "crawl_zhenai/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) Submit(request engine.Request)  {
	q.requestChan <- request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request)  {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run()  {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ [] chan engine.Request   //每一个goroutine都有一个请求队列和工作队列
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0&&len(workerQ)>0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <- q.requestChan:
				requestQ = append(requestQ, r)
			case w := <- q.workerChan:
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:   //过来了一个激活的请求，然后被激活的工作chan接受
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
