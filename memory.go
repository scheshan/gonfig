package gonfig

type memorySource struct {
	data map[string]string
}

func (s *memorySource) GetData() map[string]string {
	return s.data
}

//AddMemory add memory map to configuration
func AddMemory(builder IConfigBuilder, data map[string]string) {
	s := &memorySource{
		data: data,
	}
	builder.AddSource(s)
}
