package selector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	checkURL = "http://bt2.rutracker.org/myip?json"
)

type checkResponse struct {
	RemoteAddr string `json:"REMOTE_ADDR"`
	Proxy      string `json:"proxy"`
}

func checkProxy(t ProxyType, addr string) (*http.Client, error) {
	client, err := getClient(t, addr)
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(checkURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = compareResult(body, strings.Split(addr, ":")[0]); err != nil {
		return nil, err
	}
	return client, nil
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
