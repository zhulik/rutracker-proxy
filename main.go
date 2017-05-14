package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

var (
	port            = flag.Int("p", 8080, "Proxy port(8080)")
	rotationTimeout = flag.Int("r", 5, "Proxy rotation timeout in minutes(5 minutes)")
	proxyType       = flag.String("t", "http", "Proxy type http|socks (http)")
)

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()

	p := selector.HTTP
	switch *proxyType {
	case "http":
		break
	case "socks":
		p = selector.SOCKS
	default:
		log.Fatal("Unknown proxy type ", *proxyType)
	}

	go rotateTransport(p, proxy, (time.Duration(*rotationTimeout))*time.Minute)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy))
}
