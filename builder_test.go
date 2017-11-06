package gonfig

import "testing"

type testConfigSource struct {
}

func (p *testConfigSource) GetData() map[string]string {
	return make(map[string]string)
}


func Test_NewBuilder(t *testing.T) {
	builder := NewBuilder()
	if builder == nil {
		t.Error("builder is nil")
	}
}

func Test_AddSource(t *testing.T) {
	builder := NewBuilder()
	builder.AddSource(&testConfigSource{})
}

func Test_GetProviderList(t *testing.T){
	builder := NewBuilder()

	s := &testConfigSource{}

	builder.AddSource(s)

	pList := builder.GetSources()

	if len(pList) != 1{
		t.Error("Source's length error")
	}
	if pList[0] != s{
		t.Error("Source type error")
	}
}

func Test_Build(t *testing.T){
	builder := NewBuilder()

	s := &testConfigSource{}

	builder.AddSource(s)

	c := builder.Build()

	config, ok := c.(*config)

	if !ok{
		t.Error("Config type error")
	}

	if len(config.sList) == 0{
		t.Error("Config source length error")
	}
}