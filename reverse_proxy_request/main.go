package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// このフィールドでパスやヘッダー, ボディーの書き換えもできる
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":9001"
	}
	// リバースプロキシの実装
	rp := &httputil.ReverseProxy{
		Director: director,
	}
	server := http.Server{
		Addr: "127.0.0.1:9000",
		Handler: rp,
	}
	log.Println("Start Listening at :9000")
	log.Fatalln(server.ListenAndServe())
}
