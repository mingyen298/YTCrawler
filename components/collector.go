package ytcrawler

// import (
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"time"
// 	"yt_crawler/constants"

// 	"github.com/gocolly/colly"
// )

// const findBaseJsRegex = "(.*)base.js$"

// const TraversalBaseJSTopic = "script[nonce][src]"
// const TraversalYtInitRespTopic = "script[nonce]"

// type Extractor struct {
// 	collector         *colly.Collector
// 	extractYtInitResp *regexp.Regexp
// 	findBaseJsRegex   *regexp.Regexp
// 	content           []byte
// 	decryptRuleJson   *DecryptRuleJson
// }

// func TestSpeed(f func()) {
// 	s := time.Now()
// 	f()
// 	fmt.Println(time.Since(s))
// }

// var deviceHeaders map[string]string

// func (m *Extractor) Initial() {
// 	deviceHeaders = map[string]string{
// 		"Mac": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
// 		"LG":  "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.5249.79 Mobile Safari/537.36",
// 	}
// 	_, err := os.Stat(constants.BaseDecryptRuleJsonFile)
// 	if err != nil {
// 		return
// 	} else {

// 	}

// 	m.extractYtInitResp = regexp.MustCompile(findVideoFormatsRegex)
// 	m.collector = colly.NewCollector(
// 		colly.UserAgent(deviceHeaders["LG"]),
// 	)

// 	// m.collector.OnHTML(TraversalBaseJSTopic, m.findBaseJSURL)
// 	// m.collector.OnHTMLDetach(TraversalBaseJSTopic)
// 	m.collector.OnHTML(TraversalYtInitRespTopic, m.findYtInitialResponse)

// 	m.collector.OnRequest(func(r *colly.Request) {
// 		r.Headers.Add("Accept-Language", "en-US;q=0.8,en;q=0.9")
// 	})
// }

// func (m *Extractor) Proccess(url string) {
// 	m.collector.Visit(url)
// 	m.collector.Wait()
// }

// func (m *Extractor) findYtInitialResponse(e *colly.HTMLElement) {
// 	sub := m.extractYtInitResp.FindSubmatch([]byte(e.Text))
// 	if len(sub) != 0 {
// 		m.content = sub[1]
// 	}
// }

// func (m *Extractor) findBaseJSURL(e *colly.HTMLElement) {

// 	src := e.Attr("src")
// 	if match, _ := regexp.Match(findBaseJsRegex, []byte(src)); match {
// 		const domain = "https://www.youtube.com"
// 		newPath := domain + src
// 		c2 := m.collector.Clone()
// 		c2.OnResponse(func(r *colly.Response) {

// 			baseJSExtractor := BaseJSExtractor{}
// 			baseJSExtractor.Proccess(r.Body)

// 		})

// 		c2.Visit(newPath)

// 	}
// }

// func (m *Extractor) FindYtInitialResponse(b []byte) {
// 	sub := m.extractYtInitResp.FindSubmatch([]byte(e.Text))
// 	if len(sub) != 0 {
// 		m.content = sub[1]
// 	}
// }

// func (m *Extractor) FindBaseJSURL(b []byte) {

// 	src := e.Attr("src")
// 	if match, _ := regexp.Match(findBaseJsRegex, []byte(src)); match {
// 		const domain = "https://www.youtube.com"
// 		newPath := domain + src
// 		c2 := m.collector.Clone()
// 		c2.OnResponse(func(r *colly.Response) {

// 			baseJSExtractor := BaseJSExtractor{}
// 			baseJSExtractor.Proccess(r.Body)

// 		})

// 		c2.Visit(newPath)

// 	}
// }

// func SetTraval
