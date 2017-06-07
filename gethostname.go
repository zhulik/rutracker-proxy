// +build go1.5, go1.6, go1.7, !go1.8

package main

import (
	"net/url"
	"strings"
)

// this is workaround  block for go < 1.8
func getHostname(u *url.URL) string {
	return stripPost(u)
}

func stripPost(u *url.URL) string {
	hostport := u.Host
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]

