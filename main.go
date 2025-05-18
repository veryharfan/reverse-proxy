package main

import (
	"fmt"
	"log"
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
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	productServiceURL := os.Getenv("PRODUCT_SERVICE_URL")
	shopServiceURL := os.Getenv("SHOP_SERVICE_URL")
	orderServiceURL := os.Getenv("ORDER_SERVICE_URL")

	log.Println("User Service URL:", userServiceURL)
	log.Println("Product Service URL:", productServiceURL)
	log.Println("Shop Service URL:", shopServiceURL)
	log.Println("Order Service URL:", orderServiceURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var target string

		log.Printf("Method: %s Request URL: %s\n", r.Method, r.URL.Path)

		switch {
		case strings.HasPrefix(r.URL.Path, "/user-service"):
			target = userServiceURL
		case strings.HasPrefix(r.URL.Path, "/product-service"):
			target = productServiceURL
		case strings.HasPrefix(r.URL.Path, "/shop-service"):
			target = shopServiceURL
		case strings.HasPrefix(r.URL.Path, "/order-service"):
			target = orderServiceURL
		default:
			http.NotFound(w, r)
			return
		}

		proxy := reverseProxy(target)
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Starting reverse proxy server on :80")
	http.ListenAndServe(":80", nil)
}
