package main

import (
	"log"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

func rotateTransport(t selector.ProxyType, proxy *goproxy.ProxyHttpServer, timeout time.Duration) {
	for {
		log.Println("Rotation started...")
		transport, err := selector.GetNextProxyTransport(t)
		if err != nil {
			log.Printf("Transport rotation error: %s", err)
			continue
		}

		proxy.Tr = transport
		log.Println("Rotation finished...")
		time.Sleep(timeout)
	}
}
