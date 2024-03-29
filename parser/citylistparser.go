package parser

import (
	"crawl_zhenai/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//i:=0
	for _,c := range all {
		result.Items = append(result.Items, string(c[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			ParserFunc: ParseCity,
		})
		//i++
		//if i>=20 {
		//	break
		//}
	}
	return result
}