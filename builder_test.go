package gonfig_test

import "testing"
import "github.com/scheshan/gonfig"

func Test_NewBuilder(t *testing.T) {
	builder := gonfig.NewBuilder()
	if builder == nil {
		t.Error("builder is nil")
	}
}

func Test_AddProvider(t *testing.T) {
	builder := gonfig.NewBuilder()
	builder.AddProvider(&testConfigurationProvider{})
}

type testConfigurationProvider struct {
}

func (p *testConfigurationProvider) GetData() map[string]string {
	return make(map[string]string)
}
