package gonfig

// IConfigSource the configuration provider that can retrive configuration data
type IConfigSource interface {
	GetData() map[string]string
	SetCallbackChannel(ch chan IConfigSource)
}
