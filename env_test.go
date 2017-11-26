package gonfig

import (
	"os"
	"testing"
)

func TestEnvSource(t *testing.T) {
	k, v := randomString(10), randomString(10)

	os.Setenv(k, v)

	s := EnvSource()
	s.Load()

	_, ok := s.(*envSource)

	if !ok {
		t.Error("Should be envSource")
	}

	value, ok := s.Get(k)

	if !ok || value != v {
		t.Error("Excepted:", v, "actual:", value)
	}
}
