package engine

import (
	"log"

	"github.com/mneumi/crawler/fetcher"
	"github.com/mneumi/crawler/types"
)

type SingleEngine struct{}

func (s SingleEngine) Run(seeds ...types.Request) {
	var requests []types.Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)

		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item: %v\n", item)
		}
	}
}

func worker(r types.Request) (types.ParseResult, error) {
	log.Printf("fetching %s\n", r.Url)

	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("fetch error url: %s, err: %v", r.Url, err)
		return types.ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
