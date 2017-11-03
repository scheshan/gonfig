package gonfig

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type configIni struct {
	path string
}

const keyDelimiter = ":"

func (i *configIni) GetData() map[string]string {
	file, err := os.Open(i.path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	return i.getDataFromReader(file)
}

func (i *configIni) getDataFromReader(reader io.Reader) map[string]string {
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
			sectionPrefix = line[1:len(line)-1] + keyDelimiter
			continue
		}

		separator := strings.Index(line, "=")
		if separator < 0 {
			panic("Unrecognized line format: '" + line + "'.")
		}

		key = sectionPrefix + strings.TrimSpace(line[0:separator])
		value = strings.TrimSpace(line[separator+1:])

		//Remote quotes
		if len(value) > 0 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-2]
		}

		result[key] = value
	}

	return result
}

func AddIni(builder IConfigBuilder, path string) {
	i := &configIni{
		path: path,
	}

	builder.AddProvider(i)
}
