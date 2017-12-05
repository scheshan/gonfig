package gonfig

import (
	"os"
	"testing"
)

func TestIniSource(t *testing.T) {
	file, err := randomFile()
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file.Name())

	file.WriteString(`Key1=Value1
//comment
#comment

;comment
[Section]
Key2=Value2
Key3="Value3"`)

	s := IniSource(file.Name())
	s.Load()

	var value string
	var ok bool

	value, ok = s.Get("Key1")
	if !ok || value != "Value1" {
		t.Error("Excepted:", "Value1", "actual:", value)
	}

	value, ok = s.Get("Section:Key2")
	if !ok || value != "Value2" {
		t.Error("Excepted:", "Value2", "actual:", value)
	}

	value, ok = s.Get("Section:Key3")
	if !ok || value != "Value3" {
		t.Error("Excepted:", "Value3", "actual:", value)
	}
}

func TestIniSourceWithInvalidFormat(t *testing.T) {
	file, err := randomFile()
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file.Name())

	defer func() {
		err := recover()

		if err == nil {
			t.Error("There must be an error")
		}
	}()

	file.WriteString(`Invalid`)

	s := IniSource(file.Name())
	s.Load()
}

func TestIniWithInvalidPath(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("There must be an error")
		}
	}()

	path := randomString(20)

	s := IniSource(path)
	s.Load()
}
