package main

import (
	"log"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

func updateTransport(t selector.ProxyType, proxy *goproxy.ProxyHttpServer) {
	log.Println("Rotation started...")
	transport, err := selector.GetNextProxyTransport(t)
	if err != nil {
		log.Printf("Transport rotation error: %s", err)
		return
	}

	proxy.Tr = transport
	log.Println("Rotation finished...")
}

func rotateTransport(t selector.ProxyType, proxy *goproxy.ProxyHttpServer, timeout time.Duration) {
	for {
		if timeout == 0 {
			break
		}
		time.Sleep(timeout)
		updateTransport(t, proxy)
	}
}
