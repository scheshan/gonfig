package gonfig

// IConfigSource the configuration provider that can retrive configuration data
type IConfigSource interface {
	GetData() map[string]string
	SetCallbackChannel(ch chan IConfigSource)
}

type configSource struct {
	data map[string]string
	ch   chan IConfigSource
}

func (c *configSource) GetData() map[string]string {
	return c.data
}

func (c *configSource) SetCallbackChannel(ch chan IConfigSource) {
	c.ch = ch
}
