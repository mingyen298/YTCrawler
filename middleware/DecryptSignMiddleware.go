package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	ytcrawler "yt_crawler/components"
)

type DecryptFuncs int

const (
	FuncReverse DecryptFuncs = iota
	FuncSplice
	FuncSwap
)

type DecryptRule struct {
	Name  string
	Param int
}
type DecryptRuleJson struct {
	Rules   []DecryptRule
	FuncMap map[string]int
}

type VideoSignDecrypter struct {
	temp      []byte
	rulesJson *DecryptRuleJson
}

func (m *VideoSignDecrypter) Reload() error {
	newFile := DecryptRuleJson{}
	_, err := os.Stat(BaseDecryptRuleJsonFile)
	if err != nil {
		return err
	}

	f, err := os.Open(BaseDecryptRuleJsonFile)
	if err != nil {
		return err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &newFile)
	if err != nil {
		return err
	}
	return nil
}

func (m *VideoSignDecrypter) Proccess(input []byte) []byte {
	m.temp = input
	for _, s := range m.rulesJson.Rules {
		switch m.rulesJson.FuncMap[s.Name] {
		case int(FuncReverse):
			m.reverse()
		case int(FuncSplice):
			m.splice(s.Param)
		case int(FuncSwap):
			m.swap(s.Param)
		}
	}
	return m.temp
}

func (m *VideoSignDecrypter) swap(b int) {
	c := m.temp[0]
	length := len(m.temp)
	m.temp[0] = m.temp[b%length]
	m.temp[b%length] = c
}

func (m *VideoSignDecrypter) reverse() {
	length := len(m.temp)
	arr := make([]byte, length)
	halfLen := length / 2
	for i := 0; i <= halfLen; i++ {
		arr[i] = m.temp[length-i-1]
		arr[length-i-1] = m.temp[i]
	}
	m.temp = arr
}

func (m *VideoSignDecrypter) splice(b int) {
	m.temp = append(m.temp[0:0], m.temp[b:]...)
}

//--------------------------Middleware---------------------------

func DecryptSignMiddleware() ytcrawler.HandlerFunc {
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
