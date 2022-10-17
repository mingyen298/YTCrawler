package middleware

import (
	"encoding/json"
	"os"
	ytcrawler "yt_crawler/components"
)

const BaseDecryptRuleJsonFile = "DecryptRule.json"

//--------------------------Middleware---------------------------

func NewDecryptRuleMiddleware() ytcrawler.HandlerFunc {

	return func(c *ytcrawler.Context) {
		ruleCollector := ytcrawler.ReceiveAgent{}
		ruleCollector.Use(ExtractDecryptRuleMiddleware(), ParseDecryptRuleMiddleware())
		c1 := ruleCollector.Proccess(c.Text)
		rules := c1.PrevOutput().([]ytcrawler.DecryptRule)

		decryFuncCollector := ytcrawler.ReceiveAgent{}
		decryFuncCollector.Use(ExtractDecryptFuncMiddleware(), ParseDecryptFuncMiddleware())
		c2 := decryFuncCollector.Proccess(c.Text)
		funcMap := c2.PrevOutput().(map[string]int)

		file := ytcrawler.DecryptRuleJson{Rules: rules, FuncMap: funcMap}
		data, _ := json.Marshal(file)

		_, err := os.Stat(BaseDecryptRuleJsonFile)
		if err == nil {
			os.Remove(BaseDecryptRuleJsonFile)
		}
		f, err := os.Create(BaseDecryptRuleJsonFile)
		if err != nil {
			c.Abort()
		}
		defer f.Close()

		f.Write(data)
		c.Next(nil)
	}
}
