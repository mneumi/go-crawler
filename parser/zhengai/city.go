package zhengai

import (
	"regexp"

	"github.com/mneumi/crawler/types"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) types.ParseResult {
	re := regexp.MustCompile(cityRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := types.ParseResult{}
	for _, m := range matches {
		userName := string(m[2])
		result.Items = append(result.Items, "User: "+userName)

		result.Requests = append(result.Requests, types.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) types.ParseResult {
				return ParseProfile(contents, userName)
			},
		})
	}

	return result
}
