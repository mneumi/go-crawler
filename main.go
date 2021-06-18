package main

import (
	"github.com/mneumi/crawler/engine"
	"github.com/mneumi/crawler/parser/zhengai"
	"github.com/mneumi/crawler/scheduler"
	"github.com/mneumi/crawler/types"
)

func main() {
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(types.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: zhengai.ParseCityList,
	})
}
