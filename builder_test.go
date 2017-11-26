package gonfig

import (
	"testing"
)

func TestBuilderAdd(t *testing.T) {
	b := NewBuilder()

	s1 := &testSource{}
	s2 := &testSource{}
	d := NewFileDepend(randomString(10))

	b.Add(s1, nil)
	b.Add(s2, d)

	builder, ok := b.(*configBuilder)
	if !ok {
		t.Error("Should be configBuilder")
	}

	if len(builder.items) != 2 {
		t.Error("Builder should have 2 items")
	}

	if builder.items[0].s != s1 {
		t.Error("Should be s1")
	}
	if builder.items[0].d != nil {
		t.Error("Should be nil")
	}

	if builder.items[1].s != s2 {
		t.Error("Should be s2")
	}
	if builder.items[1].d != d {
		t.Error("Should be d")
	}
}

func TestBuilderBuild(t *testing.T) {
	b := NewBuilder()

	s1 := &testSource{}
	b.Add(s1, nil)

	c := b.Build()

	if c == nil {
		t.Error("Should not be nil")
	}

	cfg, ok := c.(*config)

	if !ok {
		t.Error("Should be config")
	}
	if len(cfg.sList) != 1 {
		t.Error("Config should have 1 item")
	}
	if cfg.sList[0] != s1 {
		t.Error("Should be s1")
	}
}
