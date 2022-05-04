package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	serverCount = 0
)

const (
	Server1 = "http://localhost:8080"
	Server2 = "http://localhost:8081"
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
}

func logRequestPayload(proxyURL string) {
	log.Printf("proxy_url: %s\n", proxyURL)
}

func getProxyURL() string {
	var servers = []string{Server1, Server2}

	server := servers[serverCount]
	serverCount++

	if serverCount >= len(servers) {
		serverCount = 0
	}

	return server
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyURL()

	logRequestPayload(url)

	serveReverseProxy(url, res, req)
}

func main() {
	var port string
	f := os.Args[1:]
	for i := 0; i < len(f); i++ {
		port = f[i]
	}
	fmt.Println("Proxe run localhost", port)
	http.HandleFunc("/", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
