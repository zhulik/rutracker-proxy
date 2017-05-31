package main

import (
	"flag"
	"log"

	"github.com/zhulik/rutracker-proxy/selector"
)

var proxyTypes = map[string]selector.ProxyType{"http": selector.HTTP, "socks": selector.SOCKS}

func main() {
	port := flag.Int("p", 8080, "Proxy port")
	rotationTimeout := flag.Int("r", 5, "Proxy rotation timeout in minutes, 0 - disabled")
	proxyType := flag.String("t", "http", "Proxy type http|socks")

	flag.Parse()

	if p, ok := proxyTypes[*proxyType]; ok {
		log.Printf("Starting proxy with port=%d type=%s rotation timeout=%d", *port, *proxyType, *rotationTimeout)
		log.Fatal(runProxy(p, *rotationTimeout, *port))
	} else {
		log.Fatal("Unknown proxy type ", *proxyType)
	}
}
