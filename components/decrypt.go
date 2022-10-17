package ytcrawler

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

type VideoDecrypterByte struct {
	temp []byte
}

func (m *VideoDecrypterByte) ProccessDynamic(input []byte, ruleJson *DecryptRuleJson) []byte {
	m.temp = input
	for _, s := range ruleJson.Rules {
		switch ruleJson.FuncMap[s.Name] {
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

func (m *VideoDecrypterByte) swap(b int) {
	c := m.temp[0]
	length := len(m.temp)
	m.temp[0] = m.temp[b%length]
	m.temp[b%length] = c
}

func (m *VideoDecrypterByte) reverse() {
	length := len(m.temp)
	arr := make([]byte, length)
	halfLen := length / 2
	for i := 0; i <= halfLen; i++ {
		arr[i] = m.temp[length-i-1]
		arr[length-i-1] = m.temp[i]
	}
	m.temp = arr
}

func (m *VideoDecrypterByte) splice(b int) {
	m.temp = append(m.temp[0:0], m.temp[b:]...)
}
