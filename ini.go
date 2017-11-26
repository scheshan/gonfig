package gonfig

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type iniSource struct {
	ConfigSource
	path string
}

func (s *iniSource) Load() {
	file, err := os.Open(s.path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := s.getDataFromReader(file)
	s.Data = data
}

func (s *iniSource) getDataFromReader(reader io.Reader) map[string]string {
	scanner := bufio.NewScanner(reader)

	var key, value, sectionPrefix string
	var result = make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()

		//Ignore blank lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		//Ignore comments
		if line[0] == ';' || line[0] == '#' || line[0] == '/' {
			continue
		}
		//[Section:header]
		if line[0] == '[' && line[len(line)-1] == ']' {
			sectionPrefix = line[1:len(line)-1] + KeyDelimiter
			continue
		}

		separator := strings.Index(line, "=")
		if separator < 0 {
			panic("Unrecognized line format: '" + line + "'.")
		}

		key = sectionPrefix + strings.TrimSpace(line[0:separator])
		value = strings.TrimSpace(line[separator+1:])

		//Remove quotes
		if len(value) > 0 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}

		result[key] = value
	}

	return result
}

func IniSource(path string) Source {
	s := &iniSource{
		path: path,
	}

	return s
}
