package gonfig

import (
	"testing"
	"strings"
)

func Test_AddIni(t *testing.T){
	path := randomString(20)

	builder := NewBuilder()
	AddIni(builder, path)

	pList := builder.GetSources()
	if len(pList) != 1{
		t.Error("Provider length error")
	}
	p, ok := pList[0].(*iniSource)
	if !ok{
		t.Error("Type cast error")
	}
	if p.path != path{
		t.Error("Path error")
	}
}

func Test_IniGetDataFromReader(t *testing.T){
	var content = `
Key1=aaa
//comment
#comment
;comment
[Section]
Key2=bbb
	`
	reader := strings.NewReader(content)

	c := &iniSource{}
	data := c.getDataFromReader(reader)

	v, ok := data["Key1"]
	if !ok{
		t.Error("Get Key1 error")
	}
	if v != "aaa"{
		t.Error("Get Key1 error")
	}

	v, ok = data["Section:Key2"]
	if !ok{
		t.Error("Get Key2 error")
	}
	if v != "bbb"{
		t.Error("Get Key2 error")
	}
}