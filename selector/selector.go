package selector

import (
	"io/ioutil"
	"net/http"
)

type ProxyType int

const (
	HTTP ProxyType = iota
	SOCKS
)

var (
	proxyTypes = map[ProxyType]string{
		HTTP:  "http://api.rufolder.net/JIkJnKmlsFIB/v2/proxies",
		SOCKS: "http://api.rufolder.net/JIkJnKmlsFIB/v2/socks"}
)

func FindProxy(t ProxyType) (string, error) {
	resp, err := http.Get(proxyTypes[t])
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if t == HTTP {
		return "http://" + string(body), nil
	}
	return string(body), nil
}
