package parser

import (
	"crawl_zhenai/engine"
	"crawl_zhenai/models"
	"github.com/bitly/go-simplejson"
	"log"
	"regexp"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		profile := parseJson(json)
		profile.Name = name
		result.Items = append(result.Items,profile)
		log.Println(profile)
	}
	return result
}

func parseJson(json []byte) models.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, err := res.Get("objectInfo").Get("basicInfo").Array()

	var profile models.Profile
	//所以我们遍历这个切片，里面使用断言来判断类型
	for k, v := range infos {
		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}

	}

	return profile
}