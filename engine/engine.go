package engine

import (
	"log"

	"github.com/mneumi/crawler/fetcher"
	"github.com/mneumi/crawler/types"
)

func Run(seeds ...types.Request) {
	var requests []types.Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("fetch error url: %s, err: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item: %v\n", item)
		}
	}
}
