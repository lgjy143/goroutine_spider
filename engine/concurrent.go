package engine

/*
工作流程，先启动本例中的Run方法，设定一个输出结果的通道，然后启动队列调度器的run方法，它回启动一个goroutine,
并初始化一个工作chan和requestChan,生成一个请求队列和工作队列，如果请求通道有值产生就加入请求队列，请求通道加入
值是靠的submit函数，即解析完以后就会加入新的请求；工作通道有值进入就加入工作队列，用的是workerReady函数，因此会根据设定的
workerCount加入一些工作队列。如果加入工作队列中的chan有request进去了，就会取消createWorker中阻塞的in（即队列中的引用），从而执行worker
 */



type Scheduler interface {
	Submit(request Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

func (c *ConcurrentEngine) Run(Seeds ...Request)  {
	out := make(chan ParseResult)
	c.Scheduler.Run()   //启动
	for i:=0;i<c.WorkerCount ;i++  {
		createWorker(c.Scheduler.WorkerChan(),out, c.Scheduler)  //共用一个out
	}
	for _, seed := range Seeds{
		c.Scheduler.Submit(seed)
	}
	for {
		result:=<-out
		for _,item:=range result.Items{
			go func() {
				c.ItemChan <- item
			}()
		}

		for _,request := range result.Requests{
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier)  {
	go func() {
		for  {
			ready.WorkerReady(in)  //给workerChan 传入一个值，每循环一次就给worker队列加一个，这个值实际上消费的是activeWorker
			request:= <-in   //等待消费activeRequest
			result, err := worker(request)
			if err!=nil {
				continue
			}
			out <- result
		}
	}()
}