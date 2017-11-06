package gonfig

type configBuilder struct {
	sList []IConfigSource
}

func (b *configBuilder) AddSource(s IConfigSource) {
	b.sList = append(b.sList, s)
}

func (b *configBuilder) GetSources() []IConfigSource {
	return b.sList
}

func (b *configBuilder) Build() IConfig {
	return newConfig(b.sList)
}

// IConfigBuilder the builder interface to build configuration instance
type IConfigBuilder interface {
	AddSource(p IConfigSource)
	GetSources() []IConfigSource
	Build() IConfig
}

// NewBuilder create a new builder instance
func NewBuilder() IConfigBuilder {
	builder := &configBuilder{}
	builder.sList = make([]IConfigSource, 0)

	return builder
}
