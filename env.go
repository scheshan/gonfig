package gonfig

import (
	"os"
	"strings"
)

type configEnv struct {
}

func (c *configEnv) GetData() map[string]string {
	data := os.Environ()
	result := make(map[string]string)

	for _, v := range data{
		key, value := c.getKeyValue(v)
		result[key]=value
	}
	return result
}

func (c *configEnv) getKeyValue(env string) (key string, value string){
	i := strings.Index(env, "=")
	key = env[0:i]
	value = env[i+1:]
	return
}

func AddEnviron(builder IConfigBuilder){
	builder.AddProvider(&configEnv{})
}