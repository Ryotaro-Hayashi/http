package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	cert, _ := ioutil.ReadFile("/etc/ssl/ca.crt")
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	//tlsConfig.BuildNameToCertificate()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://localhost:18443")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	dump, _ := httputil.DumpResponse(resp, true)
	log.Println(string(dump))
}
