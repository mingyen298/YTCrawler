package main

import (
	"io/ioutil"
	"os"
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
// func DecryptSign(fake string) string {

// 	// res, _ := url.QueryUnescape(fake)
// 	// fmt.Println(res)
// 	d := ytcrawler.VideoDecrypterByte{}
// 	real := d.ProccessDynamic(&fake)
// 	fmt.Println(real)
// 	return real
// }

func Extract(url string) {
	exc := extractor.Extractor{}
	exc.Initial()
	exc.Proccess(url)
}

// ^var (.*)={([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\){([a-zA-Z]*)|(var ([a-zA-Z]*)=)\.
//"https://rr1---sn-nvoxuu5pf5o-u2xz.googlevideo.com/videoplayback?expire=1665705145\u0026ei=WVBIY6bbCZabvcAP7v-98AE\u0026ip=61.64.208.150\u0026id=o-AJBgRw-BTDBcl38Maz_v36FJBvPFRlRK1D5sD4eKVboI\u0026itag=22\u0026source=youtube\u0026requiressl=yes\u0026mh=KC\u0026mm=31%2C29\u0026mn=sn-nvoxuu5pf5o-u2xz%2Csn-un57sne7\u0026ms=au%2Crdu\u0026mv=m\u0026mvi=1\u0026pl=20\u0026initcwndbps=732500\u0026vprv=1\u0026mime=video%2Fmp4\u0026ns=xxpf44Z1_z76hJumhxyIB8YI\u0026cnr=14\u0026ratebypass=yes\u0026dur=990.238\u0026lmt=1665589984303188\u0026mt=1665683111\u0026fvip=5\u0026fexp=24001373%2C24007246\u0026c=MWEB\u0026txp=5532434\u0026n=0WCUpXip0dW-PhlZMG\u0026sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cns%2Ccnr%2Cratebypass%2Cdur%2Clmt\u0026sig=\u0026lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps\u0026lsig=AG3C_xAwRQIgLC7ApiFi96FB4aZTXjnO6yr6eybBmvhnMFd8622d5FwCIQDD1kEvewyNrXEVVzzpJKyh_5G72yMNARLbNEoQtPZ3pg%3D%3D"
func main() {
	// os.Remove("test/test.json")
	//3OAOq0QJ8wRgIhAMFHYaTwfYOeRmT7kiTgAphZdXSGRADOcJOkJtXVtv-MAiEA-CR8iv6uu-frwIvhgIWNjQHAV=acBRES_MHGtJ_jhQw=u=u
	// url := "https://www.youtube.com/watch?v=0zhV99Bvrgg&list=RD0zhV99Bvrgg&start_radio=1"
	// url := "https://www.youtube.com/watch?v=0zhV99Bvrgg"
	// "https://www.youtube.com/watch?v=URUIcYDq3_I&list=RDURUIcYDq3_I&start_radio=1"
	// "https://www.youtube.com/watch?v=Do4Uv1CENTQ"
	// "https://www.youtube.com/watch?v=sFS14MHk9nE"
	// url := "https://www.youtube.com/watch?v=sFS14MHk9nE"

	// Extract(url)
	f, err := os.Open("test/test.js")
	if err != nil {
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	base := extractor.BaseJSExtractor{}
	// base.FindFeatures(b)
	base.Proccess(b)
	// for _, value := range re.FindSubmatch(b) {
	// 	fmt.Println(string(value))
	// }
	/*
	   1. 所有摳出來的部分
	   2. 加密的class或是你要說object
	   3.
	*/
	/*
	   var ([a-zA-Z]*)={([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},\n([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},\n([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))}};
	*/

	// loc := re.FindIndex(b)
	// fmt.Println(loc)

	// DecryptSign("3OAOq0QJ8wRgIhAMFHYaTwfYOeRmT7kiTgAphZdXSGRADOcJOkJtXVtv-MAiEA-CR8iv6uu-frwIvhgIWNjQHAV=acBRES_MHGtJ_jhQw=u=u")
}

//https://rr2---sn-nvoxuu5pf5o-u2xz.googlevideo.com/videoplayback?expire=1665709080&ei=uF9IY5DUItPVgAOB46R4&ip=61.64.208.150&id=o-AKk-Kv1uieuRwjn-SEnUCU_bLsVNtllawbJ1WOfKPlT9&itag=140&source=youtube&requiressl=yes&mh=C-&mm=31%2C29&mn=sn-nvoxuu5pf5o-u2xz%2Csn-un57enel&ms=au%2Crdu&mv=m&mvi=2&pl=20&initcwndbps=667500&vprv=1&mime=audio%2Fmp4&ns=5vknPpNutBgrJZZsks54if4I&gir=yes&clen=5229467&dur=323.082&lmt=1574703863808131&mt=1665686952&fvip=5&keepalive=yes&fexp=24001373%2C24007246&c=MWEB&txp=5531432&n=UKhemYXX6b46LXzBDl&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cns%2Cgir%2Cclen%2Cdur%2Clmt&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AG3C_xAwRQIgUcaFfNwMWhQc9SmyAR1ZXLQh8VETudnI9DlSaOkdfX0CIQDO048sXABVKbhv9vTs2pyXIOWL6HAzNezc3zx09XQhGQ%3D%3D&sig=AOq0QJ8wRgIhAMFHYaTwfYOeRmT7kiTgAphZdXSGRODOcJ3kJtXVtv-MAiEA-CR8iu6uu-frwIvhgIWNjQHAVvacBRES_MHGtJ_jhQw=

/*
var tz={Pn:function(a){a.reverse()},
Tg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
Tn:function(a,b){a.splice(0,b)}};

var tz={Pn:function(a){a.reverse()},
Tg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
Tn:function(a,b){a.splice(0,b)}};


var tz={Tg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},
Pn:function(a){a.reverse()},
Tn:function(a,b){a.splice(0,b)}};

var tz={Tg:function(a,b){var c=a[0];a[0]=a[b%a.length];a[b%a.length]=c},Pn:function(a){a.reverse()},Tn:function(a,b){a.splice(0,b)}};
*/

/*

(var ([a-zA-Z]*)={([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))}};)

(var ([a-zA-Z]*)={([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},\n([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},\n([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))}};)
*/
