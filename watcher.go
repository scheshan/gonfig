package gonfig

import (
	"sync"

	"github.com/fsnotify/fsnotify"
)

var watcher = newFileWatcher()

type fileWatcher struct {
	w    *fsnotify.Watcher
	lock *sync.Mutex
	m    map[string][]Source
}

func (w *fileWatcher) Add(path string, s Source) error {
	w.lock.Lock()
	defer w.lock.Unlock()

	if err := w.CheckInitialize(); err != nil {
		return err
	}

	w.w.Add(path)
	w.m[path] = append(w.m[path], s)
	return nil
}

func (w *fileWatcher) CheckInitialize() error {
	if w.w == nil {
		var err error
		w.w, err = fsnotify.NewWatcher()

		if err == nil {
			go w.Watch()
		}

		return err
	}

	return nil
}

func (w *fileWatcher) Watch() {
	for e := range w.w.Events {
		sList := w.m[e.Name]

		for _, s := range sList {
			s.Load()
		}
	}
}

func newFileWatcher() *fileWatcher {
	w := &fileWatcher{}
	w.lock = new(sync.Mutex)
	w.m = make(map[string][]Source)

	return w
}
