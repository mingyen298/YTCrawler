package ytcrawler

type HandlerFunc func(*Context)
type HandlersChain []HandlerFunc

type ReceiveAgent struct {
	Handlers HandlersChain
}

func (m *ReceiveAgent) Use(layers ...HandlerFunc) {
	m.Handlers = append(m.Handlers, layers...)
}

func (m *ReceiveAgent) Proccess(text []byte) *Context {
	c := &Context{handlers: m.Handlers, Text: text, index: -1}
	c.Next(text)

	return c
}
