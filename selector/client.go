package selector

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

func getTransport(t ProxyType, addr string) (*http.Transport, error) {
	transport := http.Transport{}
	switch t {
	case HTTP:
		urlProxy, err := url.Parse("http://" + addr)
		if err != nil {
			return nil, err
		}

		transport.Proxy = http.ProxyURL(urlProxy)
		break
	case SOCKS:
		dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
		if err != nil {
			return nil, err
		}
		transport.Dial = dialer.Dial
		break
	}
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &transport, nil
}
