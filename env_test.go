package gonfig

import (
	"os"
	"testing"
	"time"
	"math/rand"
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

func randomString(strlen int) string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}