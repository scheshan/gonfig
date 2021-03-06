package gonfig

import "testing"
import "os"
import "time"

func TestFileWatcherAdd(t *testing.T) {
	s := &testSource{}

	file, err := randomFile()
	defer os.Remove(file.Name())

	if err != nil {
		t.Error(err)
	}

	watcher.Add(file.Name(), s)

	file.WriteString(randomString(20))
	file.Close()

	<-time.After(100 * time.Millisecond)

	if s.loadCount != 1 {
		t.Error("Load should be triggered")
	}
}
