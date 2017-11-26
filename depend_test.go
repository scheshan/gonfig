package gonfig

import "testing"
import "time"

func TestFileDependSubscribe(t *testing.T) {
	s := &testSource{}

	file, err := randomFile()
	if err != nil {
		t.Error(err)
	}

	d := NewFileDepend(file.Name())
	d.Subscribe(s)

	file.WriteString(randomString(10))
	file.Close()

	<-time.After(100 * time.Millisecond)

	if s.loadCount != 1 {
		t.Error("Load should be triggered")
	}
}
