package gonfig

import "testing"

func Test_NewBuilder(t *testing.T) {
	builder := NewBuilder()
	if builder == nil {
		t.Error("builder is nil")
	}
}

func Test_AddProvider(t *testing.T) {
	builder := NewBuilder()
	builder.AddProvider(&testConfigProvider{})
}

func Test_GetProviderList(t *testing.T){
	builder := NewBuilder()

	p := &testConfigProvider{}

	builder.AddProvider(p)

	pList := builder.GetProviders()

	if len(pList) != 1{
		t.Error("Provider's length error")
	}
	if pList[0] != p{
		t.Error("Provider type error")
	}
}

type testConfigProvider struct {
}

func (p *testConfigProvider) GetData() map[string]string {
	return make(map[string]string)
}
