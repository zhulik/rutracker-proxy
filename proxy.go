package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/zhulik/rutracker-proxy/selector"
)

var rutrackerHostsRE = regexp.MustCompile(`^bt[2-5]?\.(rutracker\.org|t-ru\.org|rutracker\.cc)$`)

func runProxy(p selector.ProxyType, rotationTimeout int, port int) error {
	proxy := goproxy.NewProxyHttpServer()
	go rotateTransport(p, proxy, (time.Duration(rotationTimeout))*time.Minute)
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
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
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
}
