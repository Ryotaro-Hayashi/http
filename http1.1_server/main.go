package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening :18443")
	if err := http.ListenAndServeTLS(":18443", "/etc/ssl/server.crt", "/etc/ssl/server.key", nil); err != nil {
		log.Println(err)
	}
}
