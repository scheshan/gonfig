package gonfig

type Builder interface {
	Add(s Source)
	Build() Config
}

type configBuilder struct {
	sList []Source
}

func (b *configBuilder) Add(s Source) {
	b.sList = append(b.sList, s)
}

func (b *configBuilder) Build() Config {
	return newConfig(b.sList)
}

func (b *configBuilder) Init() {
	b.sList = make([]Source, 0)
}

func NewBuilder() Builder {
	b := &configBuilder{}
	b.Init()

	return b
}
