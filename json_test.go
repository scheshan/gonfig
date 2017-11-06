package gonfig

import "testing"

import "strings"

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
