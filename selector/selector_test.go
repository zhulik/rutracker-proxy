package selector_test

import (
	"testing"

	"github.com/zhulik/rutracker-proxy/selector"
)

func TestGetNextProxyTransport(t *testing.T) {
	_, err := selector.GetNextProxyTransport(selector.HTTP)
	if err != nil {
		t.Fail()
	}

	_, err = selector.GetNextProxyTransport(selector.SOCKS)
	if err != nil {
		t.Fail()
	}
}
