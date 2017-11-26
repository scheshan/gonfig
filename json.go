package gonfig

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type jsonSource struct {
	ConfigSource
	path string
}

func (s *jsonSource) Load() {
	file, err := os.Open(s.path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := s.getDataFromReader(file)
	s.Data = data
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
	if len(prefix) > 0 {
		key = prefix + KeyDelimiter + key
	}

	//OK is true means this key store a nested object
	if m, ok := value.(map[string]interface{}); ok {
		for k, v := range m {
			s.parseJSON(result, key, k, v)
		}
	} else {
		result[key] = fmt.Sprintf("%v", value)
	}
}

func JSONSource(path string) Source {
	s := &jsonSource{
		path: path,
	}

	return s
}
