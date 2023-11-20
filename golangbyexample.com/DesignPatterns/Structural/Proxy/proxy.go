package main

type proxy struct {
	app               *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NginxProxyServer() *proxy {
	return &proxy{
		app:               &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (ngnix *proxy) handleRequest(url, method string) (int, string) {
	allowed := ngnix.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return ngnix.app.handleRequest(url, method)
}

func (ngnix *proxy) checkRateLimiting(url string) bool {
	if ngnix.rateLimiter[url] == 0 {
		ngnix.rateLimiter[url] = 1
	}
	if ngnix.rateLimiter[url] > ngnix.maxAllowedRequest {
		return false
	}
	ngnix.rateLimiter[url] = ngnix.rateLimiter[url] + 1
	return true
}
