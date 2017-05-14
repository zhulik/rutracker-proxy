package main

import (
	"fmt"

	"github.com/zhulik/rutracker-proxy/selector"
)

func main() {
	HTTPProxy, err := selector.FindProxy(selector.HTTP)
	if err != nil {
		panic(err)
	}
	SOCKSProxy, err := selector.FindProxy(selector.SOCKS)
	if err != nil {
		panic(err)
	}
	fmt.Println(selector.CheckHTTPProxy(HTTPProxy))
	fmt.Println(selector.CheckSOCKSProxy(SOCKSProxy))
}
