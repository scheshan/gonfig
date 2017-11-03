package gonfig

// IConfigurationProvider the configuration provider that can retrive configuration data
type IConfigProvider interface {
	GetData() map[string]string
}
