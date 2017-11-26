package gonfig

import (
	"time"
)

type Depend interface {
	Subscribe(s Source)
}

type fileDepend struct {
	path string
}

func (d *fileDepend) Subscribe(s Source) {
	watcher.Add(d.path, s)
}

func NewFileDepend(path string) Depend {
	d := &fileDepend{
		path: path,
	}
	return d
}

type timerDepend struct {
	d     time.Duration
	sList []Source
}

func (d *timerDepend) Subscribe(s Source) {
	d.sList = append(d.sList, s)
}

func (d *timerDepend) Start() {
	for {
		<-time.After(d.d)
		for _, s := range d.sList {
			s.Load()
		}
	}
}

func NewTimerDepend(d time.Duration) Depend {
	dep := &timerDepend{
		d:     d,
		sList: make([]Source, 0),
	}
	go dep.Start()
	return dep
}
