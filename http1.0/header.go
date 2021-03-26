package main

import (
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Head("http://localhost:18888")
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
