package middleware

import (
	"errors"
	"log"
	ytcrawler "yt_crawler/components"
)

const findDecryptFuncRegex = `var ([a-zA-Z]*)={([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},\n([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))},\n([a-zA-Z]*):function\(([a-zA-Z]*|[a-zA-Z]*,[a-zA-Z]*)\)\{([a-zA-Z]*\.reverse\(\)|var [a-zA-Z]*=[a-zA-Z]*\[0\];[a-zA-Z]*\[0\]=[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\];[a-zA-Z]*\[[a-zA-Z]*%[a-zA-Z]*\.length\]=[a-zA-Z]*|[a-zA-Z]*\.splice\(0,[a-zA-Z]*\))}};`

type DecryptFuncExtractor struct {
	ExtractorBase
}

var decryptFuncExtractor DecryptFuncExtractor = DecryptFuncExtractor{}

func (m *DecryptFuncExtractor) Extraction(b []byte) ([][]byte, error) {
	sub := m.filter.FindSubmatch(b)
	if len(sub) != 0 {
		return sub, nil
	}

	return nil, errors.New("DecryptFunc not find")
}

//--------------------------Middleware---------------------------

func ExtractDecryptFuncMiddleware() ytcrawler.HandlerFunc {
	decryptFuncExtractor.Initial(findDecryptFuncRegex)
	return func(c *ytcrawler.Context) {
		decryptFuncs, err := decryptFuncExtractor.Extraction(c.Text)
		if err != nil {
			c.Abort()
			log.Println(err.Error())
		}

		c.Next(decryptFuncs)
	}
}
