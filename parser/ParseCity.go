package parser

import (
	"crawl_zhenai/engine"
	"regexp"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrlRe = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	for _,c := range all{
		result.Items = append(result.Items,"User:"+string(c[2]))
		name:=string(c[2])
		result.Requests = append(result.Requests,engine.Request{
			Url: string(c[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	//爬取下一页
	cityRe:=regexp.MustCompile(cityUrlRe)
	all2 := cityRe.FindAllSubmatch(contents, -1)
	for _, c2 := range all2{
		result.Requests = append(result.Requests, engine.Request{
			Url:string(c2[1]),
			ParserFunc:ParseCity,
		})
	}
	return result
}
