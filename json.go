package gonfig

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type jsonSource struct {
	path string
	ch chan IConfigSource
}

func (s *jsonSource) GetData() map[string]string {
	file, err := os.Open(s.path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return s.getDataFromReader(file)
}

func (s *jsonSource) SetCallbackChannel(ch chan IConfigSource){
	s.ch = ch
}

func (s *jsonSource) getDataFromReader(reader io.Reader) map[string]string {
	result := make(map[string]string)

	//Define a temp map to store unmarshaled data
	tmp := make(map[string]interface{})

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, &tmp)

	for k, v := range tmp {
		s.parseJSON(result, "", k, v)
	}

	return result
}

func (s *jsonSource) parseJSON(result map[string]string, prefix string, key string, value interface{}) {
	//OK is true means this key store a nested object
	if m, ok := value.(map[string]interface{}); ok {
		for k, v := range m {
			s.parseJSON(result, key, k, v)
			return
		}
	}

	if len(prefix) > 0 {
		key = prefix + ":" + key
	}
	result[key] = fmt.Sprintf("%v", value)
}

//AddJSON add a json file to configuration
func AddJSON(builder IConfigBuilder, path string) {
	s := &jsonSource{
		path: path,
	}
	builder.AddSource(s)
}
