package gonfig

import (
	"os"
	"testing"
)

func TestJSONSource(t *testing.T) {
	json := `
{
	"Key1":"Value1",
	"Key2":{
		"Key3": "Value3",
		"Key4": {
			"Key5": "Value5"
		},
		"Key6": "Value6"
	}
}`
	file, err := randomFile()

	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file.Name())

	file.WriteString(json)
	file.Close()

	s := JSONSource(file.Name())
	s.Load()

	var value string
	var ok bool

	value, ok = s.Get("Key1")
	if !ok || value != "Value1" {
		t.Error("Excepted:Value1, actual:", value)
	}
	value, ok = s.Get("Key2" + KeyDelimiter + "Key3")
	if !ok || value != "Value3" {
		t.Error("Excepted:Value3, actual:", value)
	}
	value, ok = s.Get("Key2" + KeyDelimiter + "Key4" + KeyDelimiter + "Key5")
	if !ok || value != "Value5" {
		t.Error("Excepted:Value5, actual:", value)
	}
	value, ok = s.Get("Key2" + KeyDelimiter + "Key6")
	if !ok || value != "Value6" {
		t.Error("Excepted:Value6, actual:", value)
	}
}
