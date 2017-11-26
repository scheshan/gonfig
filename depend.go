package gonfig

type Depend interface {
	Subscribe(s Source)
	Unsubscribe(s Source)
}

type ConfigDepend struct {
	sList []Source
}

func (d *ConfigDepend) Subscribe(s Source) {
	d.sList = append(d.sList, s)
}

func (d *ConfigDepend) Unsubscribe(s Source) {
	i := 0

	for j, source := range d.sList {
		if source == s {
			i = j
			break
		}
	}

	if i >= 0 {
		d.sList = append(d.sList[0:i], d.sList[i+1:]...)
	}
}

func (d *ConfigDepend) Init() {
	d.sList = make([]Source, 0)
}

func (d *ConfigDepend) Notify() {
	for _, s := range d.sList {
		s.Load()
	}
}
