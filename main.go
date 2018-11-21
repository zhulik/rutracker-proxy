package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/zhulik/rutracker-proxy/selector"
)

var proxyTypes = map[string]selector.ProxyType{"http": selector.HTTP, "socks": selector.SOCKS}

func main() {
	port := flag.Int("p", 8080, "Proxy port")
	rotationTimeout := flag.Int("r", 5, "Proxy rotation timeout in minutes, 0 - disabled")
	proxyType := flag.String("t", "http", "Proxy type http|socks")
	maxTries := flag.Int("m", 5, "Maximum number of requests to rufolder before giving up")

	flag.Parse()

	if p, ok := proxyTypes[*proxyType]; ok {
		log.Printf("Starting proxy with port=%d type=%s rotation timeout=%d maxTries=%d",
			*port, *proxyType, *rotationTimeout, *maxTries)
		proxy := newProxy(p, *rotationTimeout, *port, *maxTries)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy))

	} else {
		log.Fatal("Unknown proxy type ", *proxyType)
	}
}
