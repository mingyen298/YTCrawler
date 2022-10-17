package middleware

import (
	"errors"
	"log"
	"regexp"
	ytcrawler "yt_crawler/components"
)

const findBaseJsRegex = "(.*)base.js$"

type BaseJsUrlExtractor struct {
	ExtractorBase
}

var baseJsUrlExtractor BaseJsUrlExtractor = BaseJsUrlExtractor{}

func (m *BaseJsUrlExtractor) Extraction(b []byte) ([]byte, error) {

	if match, _ := regexp.Match(findBaseJsRegex, b); match {
		const domain = "https://www.youtube.com"
		newPath := domain + string(b)
		return []byte(newPath), nil
	}

	return nil, errors.New("BaseJsUrl not find")
}

//--------------------------Middleware---------------------------

func ExtractBaseJsUrlMiddleware() ytcrawler.HandlerFunc {
	baseJsUrlExtractor.Initial(findBaseJsRegex)
	return func(c *ytcrawler.Context) {
		result, err := baseJsUrlExtractor.Extraction(c.Text)
		if err != nil {
			c.Abort()
			log.Println(err.Error())
		}
		c.Next(result)
	}
}
