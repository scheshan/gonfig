package gonfig

//Builder is the interface that can configure and build a Config instance.
type Builder interface {
	Add(s Source, d Depend)
	Build() Config
}

type builderItem struct {
	s Source
	d Depend
}

type configBuilder struct {
	items []*builderItem
}

func (b *configBuilder) Add(s Source, d Depend) {
	item := &builderItem{
		s: s,
		d: d,
	}
	b.items = append(b.items, item)
}

func (b *configBuilder) Build() Config {
	return newConfig(b.items)
}

func (b *configBuilder) Init() {
	b.items = make([]*builderItem, 0)
}

//NewBuilder will return a Builder instance
func NewBuilder() Builder {
	b := &configBuilder{}
	b.Init()

	return b
}
