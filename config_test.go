package gonfig

import (
	"testing"
)

func TestConfigGetData(t *testing.T) {
	keyDuplicate := randomString(10)

	data1 := make(map[string]string)
	data1[randomString(10)] = randomString(10)
	data1[keyDuplicate] = randomString(10)

	data2 := make(map[string]string)
	data2[randomString(10)] = randomString(10)
	data2[keyDuplicate] = randomString(10)

	//builder := NewBuilder()
}
