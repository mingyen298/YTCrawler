package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	ytcrawler "yt_crawler/components"
	"yt_crawler/middleware"
)

/*
可能的難點
1. lsig= 應該是為了防護所做的字串，因為實際簽名是sig= 所以以正則表達式的觀點來看需要，需要先轉unicode才能將值取出
2. url分為兩種形式 一種塞在signatrueCipher 一種就放在url後面
3. streamingData 底下的formats、adaptiveFormats都要去逛，因為裡面都有影片
4. 要寫一個爬蟲定期更新base.js中的混淆算法，用regex做應該挺快的，然後將結果存在txt使用時將檔案讀入，如果某次解密影片打不開就可以直接呼叫這個爬蟲更新檔案
目前測試結果證明:
signCipher 不會看coockie 、device header資訊 ，IP也無所謂(電腦是家裡的網路、手機是4G下測試)
url 中的沒有成功過，也許可以試試把lsig拿去解簽名，因為到目前為止都是拿sig去解的
*/

func TestSpeed(f func()) {
	s := time.Now()
	f()
	fmt.Println(time.Since(s))
}

// func DecryptSign(fake string) string {

// 	// res, _ := url.QueryUnescape(fake)
// 	// fmt.Println(res)
// 	d := ytcrawler.VideoDecrypterByte{}
// 	real := d.Proccess(&fake)
// 	fmt.Println(real)
// 	return real
// }

func Test2() {
	fake := "Z=gXbsqCBBs3awYhAFD=9crBYVtSZGC0-YJQZtAvnV2abCQICYQ-u5ATrZMDuITNylZ9_7LqA_hCTQMx1mO1hhadDnNIgIQRw8JQ0AObAOb"
	f, err := os.Open("DecryptRule.json")
	if err != nil {
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	j := ytcrawler.DecryptRuleJson{}
	json.Unmarshal(b, &j)

	vdb := ytcrawler.VideoDecrypterByte{}

	TestSpeed(func() {
		fmt.Println(string(vdb.ProccessDynamic([]byte(fake), &j)))
	})
	fmt.Println()

	// TestSpeed(func() {
	// 	fmt.Println(string(vdb.Proccess([]byte(fake))))
	// })

	// TestSpeed(func() {
	// 	fmt.Println(string(vdb.ProccessDynamic([]byte(fake), &j)))
	// })
}

func ReadFile(name string) (b []byte) {
	f, err := os.Open(name)
	if err != nil {
		return nil
	}
	defer f.Close()
	b, err = ioutil.ReadAll(f)
	if err != nil {
		return nil
	}
	return
}

func main() {
	// fake := "Z=gXbsqCBBs3awYhAFD=9crBYVtSZGC0-YJQZtAvnV2abCQICYQ-u5ATrZMDuITNylZ9_7LqA_hCTQMx1mO1hhadDnNIgIQRw8JQ0AObAOb"

	// TestSpeed(func() {
	// 	fmt.Println(decrypt.Reverse([]byte(fake)))
	// })

	// TestSpeed(func() {
	// 	b := []byte(fake)
	// 	fmt.Println(fake)
	// 	fmt.Println(len(fake))

	// 	n := 2
	// 	temp := decrypt.Splice1(b, n)
	// 	fmt.Println(string(temp))
	// 	fmt.Println(len(string(temp)))

	// 	temp2 := decrypt.Splice2(b, n)
	// 	fmt.Println(string(temp2))
	// 	fmt.Println(len(string(temp2)))
	// })

	b := ReadFile("test/test.js")

	TestSpeed(func() {

		ruleCollector := ytcrawler.ReceiveAgent{}
		ruleCollector.Use(middleware.ExtractDecryptRuleMiddleware(), middleware.ParseDecryptRuleMiddleware())
		c1 := ruleCollector.Proccess(b)
		rules := c1.PrevOutput().([]ytcrawler.DecryptRule)
		fmt.Println(rules)

		decryFuncCollector := ytcrawler.ReceiveAgent{}
		decryFuncCollector.Use(middleware.ExtractDecryptFuncMiddleware(), middleware.ParseDecryptFuncMiddleware())
		c2 := decryFuncCollector.Proccess(b)
		funcMap := c2.PrevOutput().(map[string]int)
		fmt.Println(funcMap)

		file := ytcrawler.DecryptRuleJson{Rules: rules, FuncMap: funcMap}
		data, _ := json.Marshal(file)
		fmt.Println(string(data))
	})

	TestSpeed(func() {

		ruleCollector := ytcrawler.ReceiveAgent{}
		ruleCollector.Use(middleware.NewDecryptRuleMiddleware())
		ruleCollector.Proccess(b)

	})
	// b2 := ReadFile("test.html")
	// TestSpeed(func() {

	// 	ruleCollector := ytcrawler.ReceiveAgent{}
	// 	ruleCollector.Use(middleware.ExtractYtInitRespMiddleware())
	// 	c1 := ruleCollector.Proccess(b2)
	// 	resp := c1.PrevOutput().([]byte)
	// 	fmt.Println(string(resp))

	// })
	// fmt.Println()

}
