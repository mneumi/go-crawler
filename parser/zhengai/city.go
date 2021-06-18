package zhengai

import (
	"regexp"

	"github.com/mneumi/crawler/types"
)

var profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(
	`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai/[^"]+)"`,
)

func ParseCity(contents []byte) types.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

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

	cityUrlRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, types.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
