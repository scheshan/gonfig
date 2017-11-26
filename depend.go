package gonfig

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
