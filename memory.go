package gonfig

type memorySource struct {
	configSource
}

//AddMemory add memory map to configuration
func AddMemory(builder IConfigBuilder, data map[string]string) {
	s := &memorySource{}
	s.data = data
	builder.AddSource(s)
}
