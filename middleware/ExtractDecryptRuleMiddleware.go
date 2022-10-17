package middleware

import (
	"errors"
	"log"
	ytcrawler "yt_crawler/components"
)

const findRuleRegex = `[a-zA-Z]*=function\([a-zA-Z]*\)\{[a-zA-Z]*=[a-zA-Z]*\.split\(""\);(.*);return [a-zA-Z]*\.join\(""\)};`

type DecryptRuleExtractor struct {
	ExtractorBase
}

var decryptRuleExtractor DecryptRuleExtractor = DecryptRuleExtractor{}

func (m *DecryptRuleExtractor) Extraction(b []byte) ([]byte, error) {
	sub := m.filter.FindSubmatch(b)

	if len(sub) != 0 {
		return sub[1], nil
	}

	return nil, errors.New("DecryptRule not find")
}

//--------------------------Middleware---------------------------

func ExtractDecryptRuleMiddleware() ytcrawler.HandlerFunc {
	decryptRuleExtractor.Initial(findRuleRegex)
	return func(c *ytcrawler.Context) {
		result, err := decryptRuleExtractor.Extraction(c.Text)

		if err != nil {
			c.Abort()
			log.Println(err.Error())
		}
		c.Next(result)
	}
}
