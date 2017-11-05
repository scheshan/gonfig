package gonfig

import (
	"sync"
)

// IConfig configuration instance
type IConfig interface {
	Get(key string) (string, bool)
}

type config struct {
	pList []IConfigProvider
	data  map[string]string
	m     *sync.Mutex
}

func newConfig(pList []IConfigProvider) IConfig {
	c := &config{
		pList: pList,
		m:     new(sync.Mutex),
	}
	return c
}

func (c *config) load() {
	c.m.Lock()

	data := make(map[string]string)

	for _, p := range c.pList {
		pData := p.GetData()

		for k, v := range pData {
			data[k] = v
		}
	}
	c.data = data
	c.m.Unlock()
}

func (c *config) Get(key string) (result string, ok bool) {
	result, ok = c.data[key]
	return
}
