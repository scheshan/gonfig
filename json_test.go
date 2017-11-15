package gonfig

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_GetDataFromReader(t *testing.T) {
	str := `{"key1":1,"key2":"aaa","key3":true,"key4":{"key1":1},"key5":[1,2,3]}`
	reader := strings.NewReader(str)

	s := &jsonSource{}
	data := s.getDataFromReader(reader)

	if data["key1"] != "1" {
		t.Error("Get key1 error")
	}

	if data["key2"] != "aaa" {
		t.Error("Get key2 error")
	}

	if data["key3"] != "true" {
		t.Error("Get key3 error")
	}

	if data["key4:key1"] != "1" {
		t.Error("Get key4:key1 error")
	}

	if data["key5"] != "[1 2 3]" {
		t.Error("Get key5 error")
	}
}

func Test_AddJSON(t *testing.T) {
	path := randomString(20)

	builder := NewBuilder()
	AddJSON(builder, path)

	pList := builder.GetSources()
	if len(pList) != 1 {
		t.Error("Provider length error")
	}
	p, ok := pList[0].(*jsonSource)
	if !ok {
		t.Error("Type cast error")
	}
	if p.path != path {
		t.Error("Path error")
	}
}

func Test_JSONSetCallbackChannel(t *testing.T) {
	s := &jsonSource{}
	ch := make(chan IConfigSource)
	s.SetCallbackChannel(ch)

	if s.ch != ch {
		t.Error("Set callback channel error")
	}
}

func Test_JSONGetData(t *testing.T) {
	p := filepath.Join(os.TempDir(), randomString(20))

	fmt.Println(p)

	defer os.Remove(p)

	file, err := os.Create(p)
	if err != nil {
		t.Error(err)
	}

	content := `{"Key1":"Value1"}`

	file.WriteString(content)
	file.Close()

	c := &jsonSource{
		path: p,
	}
	data := c.GetData()

	v, ok := data["Key1"]
	if !ok {
		t.Error("Get Key1 error")
	}
	if v != "Value1" {
		t.Error("Get Key1 error")
	}
}
