package gonfig

const KeyDelimiter = ":"

type Config interface {
	Get(key string) (string, bool)
	Map(i interface{})
}

type config struct {
	sList []Source
}

func (c *config) Get(key string) (value string, ok bool) {
	for i := len(c.sList); i > 0; i-- {
		if value, ok = c.sList[i-1].Get(key); ok {
			return
		}
	}

	return
}

func (c *config) Map(i interface{}) {

}

func newConfig(items []*builderItem) *config {
	sList := make([]Source, len(items))
	for i := range items {
		s, d := items[i].s, items[i].d
		s.Load()
		if d != nil {
			d.Subscribe(s)
		}

		sList[i] = s
	}

	c := &config{
		sList: sList,
	}

	return c
}
