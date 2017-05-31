package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

func runProxy(p selector.ProxyType, rotationTimeout int, port int) error {
	proxy := goproxy.NewProxyHttpServer()
	go rotateTransport(p, proxy, (time.Duration(rotationTimeout))*time.Minute)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
}
