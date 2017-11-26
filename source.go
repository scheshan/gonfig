package gonfig

type Source interface {
	Load()
	Get(key string) (value string, ok bool)
}

type ConfigSource struct {
	Data map[string]string
}

func (s *ConfigSource) Load() {
	s.Data = make(map[string]string)
}

func (s *ConfigSource) Get(key string) (value string, ok bool) {
	value, ok = s.Data[key]

	return
}
