package main

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

var rutrackerHostsRE = regexp.MustCompile(`^bt[2-5]?\.(rutracker\.org|t-ru\.org|rutracker\.cc)$`)

func proxyHandler(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if rutrackerHostsRE.MatchString(req.URL.Hostname()) {
		log.Printf("Querying to %s through proxy...", req.URL)
		resp, err := ctx.RoundTrip(req)
		if err != nil {
			log.Printf("Error when requesting url through proxy %s: %s", req.URL, err.Error())
		}
		return req, resp
	}
	log.Printf("Querying to %s directly...", req.URL)
	req.RequestURI = ""
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error when requesting url directly %s: %s", req.URL, err.Error())
	}
	return req, resp
}

func newProxy(p selector.ProxyType, rotationTimeout int, port int) *goproxy.ProxyHttpServer {
	proxy := goproxy.NewProxyHttpServer()
	updateTransport(p, proxy)
	proxy.OnRequest().DoFunc(proxyHandler)
	go rotateTransport(p, proxy, (time.Duration(rotationTimeout))*time.Minute)
	return proxy
}
