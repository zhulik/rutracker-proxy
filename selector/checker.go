package selector

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

const (
	checkURL = "http://bt2.rutracker.org/myip?json"
)

type checkResponse struct {
	RemoteAddr string `json:"REMOTE_ADDR"`
	Proxy      string `json:"proxy"`
}

func CheckHTTPProxy(addr string) error {
	t := url.URL{}
	urlProxy, err := t.Parse(addr)
	if err != nil {
		return err
	}

	transport := http.Transport{}
	transport.Proxy = http.ProxyURL(urlProxy)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: &transport}
	resp, err := client.Get(checkURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return compareResult(body, urlProxy.Hostname())
}

func CheckSOCKSProxy(addr string) error {
	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	if err != nil {
		return err
	}

	transport := http.Transport{}
	// set our socks5 as the dialer
	transport.Dial = dialer.Dial
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: &transport}
	resp, err := client.Get(checkURL) // do request through proxy
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return compareResult(body, strings.Split(addr, ":")[0])
}

func compareResult(body []byte, addr string) error {
	r := checkResponse{}
	err := json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	if r.Proxy == r.RemoteAddr || r.Proxy != addr {
		return fmt.Errorf("Wrong response from %s: %s", addr, string(body))
	}
	return nil
}
