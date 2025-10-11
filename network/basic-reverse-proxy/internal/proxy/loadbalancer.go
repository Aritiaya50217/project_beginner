package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
	"time"
)

type LoadBalancer struct {
	targets []*url.URL
	counter uint64
}

func NewLoadBalancedReverseProxy(backends []string) http.Handler {
	targets := make([]*url.URL, 0)
	for _, backend := range backends {
		u, err := url.Parse(backend)
		if err != nil {
			log.Fatalf("❌ Invalid backend URL: %v", err)
		}
		targets = append(targets, u)
	}

	lb := &LoadBalancer{targets: targets}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := lb.nextTarget()
		proxy := httputil.NewSingleHostReverseProxy(target)

		start := time.Now()

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
			req.Header.Set("X-Forwarded-Host", r.Host)
			req.Header.Set("X-Origin-Host", target.Host)
		}

		proxy.ModifyResponse = func(resp *http.Response) error {
			duration := time.Since(start)
			log.Printf("✅ [%s] %s %s → %s [%d] in %v",
				r.Method, r.Host, r.URL.Path, target.Host, resp.StatusCode, duration)
			return nil
		}

		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			log.Printf("🔥 Proxy error to %s: %v", target.Host, err)
			http.Error(rw, "Proxy error: "+err.Error(), http.StatusBadGateway)
		}

		proxy.ServeHTTP(w, r)
	})
}

// nextTarget เลือก backend ถัดไปแบบ Round Robin
func (lb *LoadBalancer) nextTarget() *url.URL {
	index := atomic.AddUint64(&lb.counter, 1)
	return lb.targets[int(index)%len(lb.targets)]
}
