package main

import (
	"github.com/mneumi/crawler/engine"
	"github.com/mneumi/crawler/parser/zhengai"
	"github.com/mneumi/crawler/types"
)

func main() {
	engine.Run(types.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: zhengai.ParseCityList,
	})
}
