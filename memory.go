package gonfig

type memorySource struct {
	data map[string]string
	ch chan IConfigSource
}

func (s *memorySource) GetData() map[string]string {
	return s.data
}

func (s *memorySource) SetCallbackChannel(ch chan IConfigSource){
	s.ch = ch
}

//AddMemory add memory map to configuration
func AddMemory(builder IConfigBuilder, data map[string]string) {
	s := &memorySource{
		data: data,
	}
	builder.AddSource(s)
}
