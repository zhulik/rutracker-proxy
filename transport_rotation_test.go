package main

import (
	"testing"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

func TestUpdateTransport(t *testing.T) {
	proxy := goproxy.NewProxyHttpServer()
	tr := proxy.Tr
	updateTransport(selector.HTTP, proxy, 5)
	if proxy.Tr == tr {
		t.Fail()
	}
}
