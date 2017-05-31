package main

import (
	"testing"

	"github.com/zhulik/rutracker-proxy/selector"
)

func TestNewProxy(t *testing.T) {
	res := newProxy(selector.HTTP, 5, 8080)
	if res == nil {
		t.Fail()
	}
}
