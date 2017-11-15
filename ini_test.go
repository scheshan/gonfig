package gonfig

import (
	"strings"
	"testing"
)

const iniContent = `
Key1=aaa
//comment
#comment
;comment
[Section]
Key2=bbb
Key3="ccc"
`

func Test_AddIni(t *testing.T) {
	path := randomString(20)

	builder := NewBuilder()
	AddIni(builder, path)

	pList := builder.GetSources()
	if len(pList) != 1 {
		t.Error("Provider length error")
	}
	p, ok := pList[0].(*iniSource)
	if !ok {
		t.Error("Type cast error")
	}
	if p.path != path {
		t.Error("Path error")
	}
}

func Test_IniGetDataFromReader(t *testing.T) {
	reader := strings.NewReader(iniContent)

	c := &iniSource{}
	data := c.getDataFromReader(reader)

	v, ok := data["Key1"]
	if !ok {
		t.Error("Get Key1 error")
	}
	if v != "aaa" {
		t.Error("Get Key1 error")
	}

	v, ok = data["Section:Key2"]
	if !ok {
		t.Error("Get Key2 error")
	}
	if v != "bbb" {
		t.Error("Get Key2 error")
	}

	v, ok = data["Section:Key3"]
	if !ok {
		t.Error("Get Key3 error")
	}
	if v != "ccc" {
		t.Error("Get Key3 error")
	}
}

func Test_IniGetDataFromReaderWithInvalidFormat(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("An error should be occurred")
		}
	}()

	content := `
invalid format	
`
	reader := strings.NewReader(content)

	c := &iniSource{}
	c.getDataFromReader(reader)
}

func Test_IniSetCallbaclChannel(t *testing.T) {
	ch := make(chan IConfigSource)
	c := &iniSource{}
	c.SetCallbackChannel(ch)

	if c.ch != ch {
		t.Error("Set callback channel error")
	}
}
