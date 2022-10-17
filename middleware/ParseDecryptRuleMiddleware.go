package middleware

import (
	"errors"
	"log"
	"strconv"
	"strings"
	ytcrawler "yt_crawler/components"
)

const exctractFuncAndParamRegex = `^[a-zA-Z]*\.([a-zA-Z]*)\([a-zA-Z]*,([0-9]{1,2})\)`

type DecryptRuleParser struct {
	ExtractorBase
}

var decryptRuleParser DecryptRuleParser = DecryptRuleParser{}

func (m *DecryptRuleParser) Parse(b []byte) ([]ytcrawler.DecryptRule, error) {

	useFuncs := strings.Split(string(b), ";")
	rules := []ytcrawler.DecryptRule{}
	for _, function := range useFuncs {
		result := m.filter.FindStringSubmatch(function)

		param, _ := strconv.Atoi(result[2])
		rules = append(rules, ytcrawler.DecryptRule{Name: result[1], Param: param})

	}
	if len(rules) != 0 {
		return rules, nil
	}

	return nil, errors.New("Parse decrypt rule failed")
}

//--------------------------Middleware---------------------------

func ParseDecryptRuleMiddleware() ytcrawler.HandlerFunc {
	decryptRuleParser.Initial(exctractFuncAndParamRegex)
	return func(c *ytcrawler.Context) {
		rules, err := decryptRuleParser.Parse(c.PrevOutput().([]byte))
		if err != nil {
			c.Abort()
			log.Println(err.Error())
		}

		c.Next(rules)
	}
}
