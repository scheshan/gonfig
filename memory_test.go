package gonfig

import (
	"testing"
)

func TestMemorySource(t *testing.T) {
	data := make(map[string]string)
	k1, v1 := randomString(10), randomString(10)
	data[k1] = v1

	s := MemorySource(data)

	_, ok := s.(*memorySource)

	if !ok {
		t.Error("Should be memorySource")
	}

	s.Load()

	value1, ok := s.Get(k1)

	if !ok || value1 != v1 {
		t.Error("Excepted:", v1, "actual:", value1)
	}
}

func TestMemoryGetData(t *testing.T) {
	data := make(map[string]string)
	k1, k2, v1, v2 := randomString(10), randomString(10), randomString(10), randomString(10)
	data[k1], data[k2] = v1, v2

	s := MemorySource(data)

	var value1, value2 string
	var ok bool

	value1, ok = s.Get(k1)
	if !ok || value1 != v1 {
		t.Error("Excepted", v1, "actual", value1)
	}

	value2, ok = s.Get(k2)
	if !ok || value2 != v2 {
		t.Error("Excepted", v2, "actual", value2)
	}

	keyNotExist := randomString(20)

	_, ok = s.Get(keyNotExist)
	if ok {
		t.Error("Should be false")
	}
}
