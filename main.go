package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func reverseProxy(target string) http.Handler {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var target string

		switch {
		case strings.HasPrefix(r.URL.Path, "/user-service"):
			target = os.Getenv("USER_SERVICE_URL")
		case strings.HasPrefix(r.URL.Path, "/product-service"):
			target = os.Getenv("PRODUCT_SERVICE_URL")
		case strings.HasPrefix(r.URL.Path, "/shop-service"):
			target = os.Getenv("SHOP_SERVICE_URL")
		case strings.HasPrefix(r.URL.Path, "/order-service"):
			target = os.Getenv("ORDER_SERVICE_URL")
		default:
			http.NotFound(w, r)
			return
		}

		proxy := reverseProxy(target)
		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":80", nil)
}
