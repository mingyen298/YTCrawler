package middleware

import (
	"errors"
	"log"
	"regexp"
	ytcrawler "yt_crawler/components"
)

type ExtractorBase struct {
	started bool
	filter  *regexp.Regexp
}

func (m *ExtractorBase) Initial(regexString string) {
	if m.started {
		return
	}
	m.started = true
	m.filter = regexp.MustCompile(regexString)
}

const findVideoFormatsRegex = `var ytInitialPlayerResponse = (\{.*);`

type YtInitRespExtractor struct {
	ExtractorBase
}

var ytInitRespExtractor YtInitRespExtractor = YtInitRespExtractor{}

func (m *YtInitRespExtractor) Extraction(b []byte) ([]byte, error) {
	sub := m.filter.FindSubmatch(b)
	if len(sub) != 0 {
		return sub[1], nil
	}
	return nil, errors.New("ytInitialPlayerResponse not find")
}

//--------------------------Middleware---------------------------

func ExtractYtInitRespMiddleware() ytcrawler.HandlerFunc {
	ytInitRespExtractor.Initial(findVideoFormatsRegex)
	return func(c *ytcrawler.Context) {
		result, err := ytInitRespExtractor.Extraction(c.Text)
		if err != nil {
			c.Abort()
			log.Println(err.Error())
		}
		c.Next(result)
	}
}
