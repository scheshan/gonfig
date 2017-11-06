package gonfig

import "testing"

import "strings"

func Test_GetDataFromReader(t *testing.T) {
	str := `{"id":1,"name":"aaa","data":{"id":1}}`
	reader := strings.NewReader(str)

	s := &jsonSource{}
	data := s.getDataFromReader(reader)

	if data["id"] != "1" {
		t.Error("Get id error")
	}

	if data["name"] != "aaa" {
		t.Error("Get name error")
	}

	if data["data:id"] != "1" {
		t.Error("Get data:id error")
	}
}
