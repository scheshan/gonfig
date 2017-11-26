package gonfig

import "testing"

func TestConfigSourceLoad(t *testing.T) {
	cs := &ConfigSource{}
	cs.Load()

	if cs.Data == nil {
		t.Error("Data should not be nil")
	}

	k1, v1 := randomString(10), randomString(10)

	cs.Data[k1] = v1

	cs.Load()

	_, ok := cs.Data[k1]

	if ok {
		t.Error("Data should be another instance")
	}
}
