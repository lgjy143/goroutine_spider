package main

import (
	"crawl_zhenai/engine"
	"crawl_zhenai/parser"
	"crawl_zhenai/persist"
	"crawl_zhenai/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},   //可以在这里切换调度器
		WorkerCount:100,
		ItemChan:persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:url,
		ParserFunc:parser.ParseCityList,
	})
}
