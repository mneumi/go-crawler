package zhengai

import (
	"regexp"

	"github.com/mneumi/crawler/types"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) types.ParseResult {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)

	result := types.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City: "+string(m[2]))

		result.Requests = append(result.Requests, types.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
