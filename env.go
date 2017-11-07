package gonfig

import (
	"os"
	"strings"
)

type envSource struct {
	ch chan IConfigSource
}

func (s *envSource) GetData() map[string]string {
	data := os.Environ()
	result := make(map[string]string)

	for _, v := range data {
		key, value := s.getKeyValue(v)
		result[key] = value
	}
	return result
}

func (s *envSource) getKeyValue(env string) (key string, value string) {
	i := strings.Index(env, "=")
	key = env[0:i]
	value = env[i+1:]
	return
}

func (s *envSource) SetCallbackChannel(ch chan IConfigSource){
	s.ch = ch
}

// AddEnviron add environment variables to configuratio
func AddEnviron(builder IConfigBuilder) {
	builder.AddSource(&envSource{})
}
