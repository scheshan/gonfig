package gonfig

type Config interface {
	Get(key string) (string, bool)
	Map(i interface{})
}

type config struct {
	sList []Source
}

func (c *config) Get(key string) (value string, ok bool) {
	for i := len(c.sList); i >= 0; i-- {
		if value, ok = c.sList[i].Get(key); ok {
			return
		}
	}

	return
}

func (c *config) Map(i interface{}) {

}

func (c *config) Init() {
	for _, s := range c.sList {
		s.Load()
	}
}

func newConfig(sList []Source) *config {
	c := &config{
		sList: sList,
	}
	c.Init()

	return c
}
