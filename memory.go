package gonfig

type memorySource struct {
	ConfigSource
}

func (s *memorySource) Load() {
	//Do nothing
}

func MemorySource(data map[string]string) Source {
	s := &memorySource{}
	s.Data = data
	return s
}
