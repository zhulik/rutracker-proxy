package selector_test

import (
	"testing"

	"github.com/zhulik/rutracker-proxy/selector"
)

func TestGetNextProxyTransport(t *testing.T) {
	_, err := selector.GetNextProxyTransport(selector.HTTP, 5)
	if err != nil {
		t.Fail()
	}

	_, err = selector.GetNextProxyTransport(selector.SOCKS, 5)
	if err != nil {
		t.Fail()
	}
}
