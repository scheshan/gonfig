package gonfig

type Source interface {
	Load()
	Get(key string) (value string, ok bool)
	Depend(d Depend)
}

type ConfigSource struct {
	Data map[string]string
	D    Depend
}

func (s *ConfigSource) Load() {

}

func (s *ConfigSource) Get(key string) (value string, ok bool) {
	value, ok = s.Data[key]

	return
}

func (s *ConfigSource) Depend(d Depend) {
	if s.D != nil {
		s.D.Unsubscribe(s)
	}

	s.D = d
	s.D.Subscribe(s)
}
