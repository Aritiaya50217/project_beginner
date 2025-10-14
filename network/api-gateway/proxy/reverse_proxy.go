package proxy

import (
	"net/http/httputil"
	"net/url"
)

func NewSingleHostReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}
