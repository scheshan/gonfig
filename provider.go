package gonfig

// IConfigurationProvider the configuration provider that can retrive configuration data
type IConfigurationProvider interface {
	GetData() map[string]string
}
