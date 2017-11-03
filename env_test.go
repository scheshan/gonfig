package gonfig

import (
	"os"
	"testing"
)

func Test_EnvGetData(t *testing.T){
	key := randomString(20)
	value := randomString(100)

	os.Setenv(key, value)

	env := configEnv{}
	d := env.GetData()

	v, ok := d[key]

	if !ok{
		t.Error("should be true")
	}
	if v != value{
		t.Error("should be " + value)
	}
}

func Test_AddEnviron(t *testing.T){
	builder := NewBuilder()
	AddEnviron(builder)

	pList := builder.GetProviders()

	if len(pList) != 1{
		t.Error("Provider length error")
	}

	_, ok := pList[0].(*configEnv)
	if !ok{
		t.Error("Type cast error")
	}
}