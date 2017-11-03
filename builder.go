package gonfig

// ConfigurationBuilder the builder to build configuration instance
type configBuilder struct {
	pList []IConfigProvider
}

func (b *configBuilder) AddProvider(p IConfigProvider) {
	b.pList = append(b.pList, p)
}

func (b *configBuilder) GetProviders() []IConfigProvider{
	return b.pList
}

// IConfigBuilder the builder interface to build configuration instance
type IConfigBuilder interface {
	AddProvider(p IConfigProvider)
	GetProviders() []IConfigProvider
}

// NewBuilder create a new builder instance
func NewBuilder() IConfigBuilder {
	builder := &configBuilder{}
	builder.pList = make([]IConfigProvider, 0)

	return builder
}
