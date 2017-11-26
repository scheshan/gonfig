package gonfig

import "testing"
import "time"

func TestFileDepend(t *testing.T) {
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

func TestTimerDepend(t *testing.T) {
	s := &testSource{}

	d := NewTimerDepend(20 * time.Millisecond)
	d.Subscribe(s)

	<-time.After(105 * time.Millisecond) //wait more 5 milliseconds

	if s.loadCount != 5 {
		t.Error("Load should be triggered")
	}
}
