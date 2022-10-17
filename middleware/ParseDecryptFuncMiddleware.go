package middleware

import (
	"regexp"
	ytcrawler "yt_crawler/components"
)

const reverseFuncRegex = `^[a-zA-Z]*\.reverse\(\)$`
const swapFuncRegex = `^var [a-zA-Z]*=[a-zA-Z]*\[0\];`
const spliceFuncRegex = `^[a-zA-Z]*\.splice\(0,[a-zA-Z]*\)$`

type DecryptFuncParser struct {
	ExtractorBase
	reverseFilter *regexp.Regexp
	spliceFilter  *regexp.Regexp
	swapFilter    *regexp.Regexp
}

var decryptFuncParser DecryptFuncParser = DecryptFuncParser{}

func (m *DecryptFuncParser) Initial(regexStrings ...string) {
	if m.started {
		return
	}
	m.started = true

	m.reverseFilter = regexp.MustCompile(regexStrings[0])
	m.swapFilter = regexp.MustCompile(regexStrings[1])
	m.spliceFilter = regexp.MustCompile(regexStrings[2])
}

func (m *DecryptFuncParser) Parse(sub [][]byte) map[string]int {
	funcMap := make(map[string]int)
	for i := 0; i < 3; i++ {

		method := sub[2+(i*3)]
		feature := sub[4+(i*3)]

		if match := m.reverseFilter.Match(feature); match {
			funcMap[string(method)] = int(ytcrawler.FuncReverse)
			continue
		}
		if match := m.swapFilter.Match(feature); match {
			funcMap[string(method)] = int(ytcrawler.FuncSwap)
			continue
		}
		if match := m.spliceFilter.Match(feature); match {
			funcMap[string(method)] = int(ytcrawler.FuncSplice)
			continue

		}

	}

	return funcMap
}

//--------------------------Middleware---------------------------

func ParseDecryptFuncMiddleware() ytcrawler.HandlerFunc {
	decryptFuncParser.Initial(reverseFuncRegex, swapFuncRegex, spliceFuncRegex)
	return func(c *ytcrawler.Context) {
		funcMap := decryptFuncParser.Parse(c.PrevOutput().([][]byte))

		c.Next(funcMap)
	}
}
