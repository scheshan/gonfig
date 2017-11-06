package gonfig

import "testing"

func Test_AddMemory(t *testing.T){
	builder := NewBuilder()

	data := make(map[string]string)

	AddMemory(builder, data)

	sList := builder.GetSources()

	if len(sList) != 1{
		t.Error("Source length error")
	}

	_, ok := sList[0].(*memorySource)

	if !ok{
		t.Error("Source type error")
	}
}

func Test_MemoryGetData(t *testing.T){
	data := make(map[string]string)

	key, value := randomString(20), randomString(20)
	data[key] = value

	source := &memorySource{
		data: data,
	}

	v := source.GetData()[key]

	if v != value{
		t.Error("Get data error")
	}
}