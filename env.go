package gonfig

import (
	"os"
	"strings"
)

type envSource struct {
	ConfigSource
}

func (s *envSource) Load() {
	env := os.Environ()
	data := make(map[string]string)

	for _, v := range env {
		key, value := s.getKeyValue(v)
		data[key] = value
	}
	s.Data = data
}

func (s *envSource) getKeyValue(env string) (key string, value string) {
	i := strings.Index(env, "=")
	key = env[0:i]
	value = env[i+1:]
	return
}

func EnvSource() Source {
	return &envSource{}
}
