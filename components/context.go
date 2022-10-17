package ytcrawler

import (
	"math"
	"sync"
)

// abortIndex represents a typical value used in abort functions.
const abortIndex int8 = math.MaxInt8 >> 1

type OutputData struct {
	data interface{}
	mu   sync.RWMutex
}

func (m *OutputData) Set(data interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = data
}
func (m *OutputData) Get() interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data
}

type Context struct {
	Text     []byte
	index    int8
	handlers HandlersChain
	// This mutex protects Keys map.
	mu sync.RWMutex

	Keys       map[string]interface{}
	outputData OutputData
}

func (c *Context) Reset() {
	c.index = -1
	c.outputData.Set(nil)
	c.Keys = map[string]interface{}{}
	c.Text = nil
}

func (c *Context) Next(handoverData interface{}) {
	c.outputData.Set(handoverData)
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) PrevOutput() interface{} {
	return c.outputData.Get()
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.Keys[key]
	return
}

func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetByteSlice(key string) (bs []byte) {
	if val, ok := c.Get(key); ok && val != nil {
		bs, _ = val.([]byte)
	}
	return
}
