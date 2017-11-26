package gonfig

import (
	"testing"
	"time"
)

func TestConfigGetData(t *testing.T) {
	k1, k2, keyDuplicate := randomString(10), randomString(10), randomString(10)
	v1, v2, valueDuplicate1, valueDuplicate2 := randomString(10), randomString(10), randomString(10), randomString(10)
	keyNotExist := randomString(20)

	data1 := make(map[string]string)
	data1[k1] = v1
	data1[keyDuplicate] = valueDuplicate1
	s1 := MemorySource(data1)

	data2 := make(map[string]string)
	data2[k2] = v2
	data2[keyDuplicate] = valueDuplicate2
	s2 := MemorySource(data2)

	d := NewTimerDepend(20 * time.Millisecond).(*timerDepend)
	s3 := &testSource{}
	s4 := &testSource{}

	builder := NewBuilder()
	builder.Add(s1, nil)
	builder.Add(s2, nil)
	builder.Add(s3, nil)
	builder.Add(s4, d)

	conf := builder.Build()

	var value string
	var ok bool

	value, ok = conf.Get(k1)
	if !ok || value != v1 {
		t.Error("Expected:", v1, "actual:", value)
	}

	value, ok = conf.Get(k2)
	if !ok || value != v2 {
		t.Error("Expected:", v2, "actual:", value)
	}

	value, ok = conf.Get(keyDuplicate)
	if !ok || value != valueDuplicate2 {
		t.Error("Expected:", valueDuplicate2, "actual:", value)
	}

	value, ok = conf.Get(keyNotExist)
	if ok {
		t.Error("Key should not exist")
	}

	<-time.After(d.d + 10*time.Millisecond)

	if s3.loadCount != 1 {
		t.Error("s3 should load once")
	}
	if s4.loadCount != 2 {
		t.Error("s4 should load twice")
	}
}
