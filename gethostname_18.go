// +build go1.8

package main

import (
	"net/url"
)

// this is workaround  block for go < 1.8
func getHostname(u *url.URL) string {
	return u.Hostname()
}
