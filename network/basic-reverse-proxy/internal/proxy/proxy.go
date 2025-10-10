package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func NewReverseProxy(target string) http.Handler {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL : %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Origin-Host", targetURL.Host)
		req.Header.Set("X-Proxy-Timestamp", time.Now().Format(time.RFC3339))
		log.Printf("Forwarding %s %s to %s", req.Method, req.URL.Path, targetURL)
	}

	proxy.ModifyResponse = func(r *http.Response) error {
		log.Printf("Response from %s: %d", targetURL.Host, r.StatusCode)
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v", err)
		http.Error(w, "Proxy error : "+err.Error(), http.StatusBadGateway)
	}
	return proxy
}
