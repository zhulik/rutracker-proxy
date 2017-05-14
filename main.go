package main

import (
	"log"

	"github.com/zhulik/rutracker-proxy/selector"
)

func main() {
	HTTPProxyClient, err := selector.GetNextProxyClient(selector.HTTP)
	if err != nil {
		panic(err)
	}
	log.Println(HTTPProxyClient)
	SOCKSProxyClient, err := selector.GetNextProxyClient(selector.SOCKS)
	if err != nil {
		panic(err)
	}
	log.Println(SOCKSProxyClient)
}
