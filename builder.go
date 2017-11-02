package gonfig

// ConfigurationBuilder the builder to build configuration instance
type configurationBuilder struct {
	pList []IConfigurationProvider
}

func (b *configurationBuilder) AddProvider(p IConfigurationProvider) {
	b.pList = append(b.pList, p)
}

// IConfigurationBuilder the builder interface to build configuration instance
type IConfigurationBuilder interface {
	AddProvider(p IConfigurationProvider)
}

// NewBuilder create a new builder instance
func NewBuilder() IConfigurationBuilder {
	builder := &configurationBuilder{}
	builder.pList = make([]IConfigurationProvider, 0)

	return builder
}
