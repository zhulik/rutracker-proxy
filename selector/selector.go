package selector

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ProxyType proxy type
type ProxyType int

const (
	// HTTP proxy type
	HTTP ProxyType = iota
	// SOCKS proxy type
	SOCKS
)

var (
	proxyTypes = map[ProxyType]string{
		HTTP:  "http://api.rufolder.net/JIkJnKmlsFIB/v2/proxies",
		SOCKS: "http://api.rufolder.net/JIkJnKmlsFIB/v2/socks"}
)

func getNextProxyURL(t ProxyType) (string, error) {
	resp, err := http.Get(proxyTypes[t])
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// GetNextProxyClient returns ready http.Client with configured transport
func GetNextProxyTransport(t ProxyType) (*http.Transport, error) {
	c := 0
	for c < 5 {
		c++
		addr, err := getNextProxyURL(t)
		if err != nil {
			log.Printf("Error from getNextProxyURL: %s", err.Error())
			continue
		}
		log.Printf("Checking proxy %s...", addr)
		t, err := checkProxy(t, addr)
		if err != nil {
			log.Printf("Proxy check failed: %s", err.Error())
			continue
		}
		log.Printf("Using proxy %s", addr)
		return t, nil
	}
	return nil, fmt.Errorf("Cannot find working proxy")
}
