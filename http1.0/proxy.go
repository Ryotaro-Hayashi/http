package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyUrl, _ := url.Parse("http://localhost:18888")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	resp, _ := client.Get("http://github.com")
	dump, _ := httputil.DumpResponse(resp, true)
	log.Println(string(dump))
}
