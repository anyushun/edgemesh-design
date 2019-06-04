package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type ClientConfig struct {
	CA string
	Cert string
	Key string
	Timeout time.Duration
}

type RequestVar struct {
	Method string
	URL string
	Body string
}

var conf *ClientConfig = &ClientConfig{}
var req *RequestVar = &RequestVar{}

func main() {
	flag.StringVar(&conf.CA, "ca", "", "ca for verify server")
	flag.StringVar(&conf.Cert, "cert", "", "cert for tls")
	flag.StringVar(&conf.Key, "key", "", "key for tls")
	flag.DurationVar(&conf.Timeout, "timeout", 5*time.Second, "request timeout")
	flag.StringVar(&req.Method, "method", http.MethodGet, "request method")
	flag.StringVar(&req.URL, "url", "http://localhost:8080", "request url")
	flag.StringVar(&req.Body, "body", "", "request body")
	flag.Parse()

	tlsConfig := &tls.Config{}

	if conf.CA == "" {
		tlsConfig.InsecureSkipVerify = true
	} else {
		caData, err := ioutil.ReadFile(conf.CA)
		if err != nil {
			log.Panic(err)
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caData)
		tlsConfig.RootCAs = pool
		tlsConfig.InsecureSkipVerify = false
	}

	if conf.Cert != "" && conf.Key != "" {
		pair, err := tls.LoadX509KeyPair(conf.Cert, conf.Key)
		if err != nil {
			log.Panicln(err)
		}
		tlsConfig.Certificates = []tls.Certificate{pair}
	}

	client := &http.Client{
		Timeout:conf.Timeout,
		Transport:&http.Transport{
			TLSClientConfig:tlsConfig,
		},
	}

	request, err := http.NewRequest(req.Method, req.URL, strings.NewReader(req.Body))
	if err != nil {
		log.Panicln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	log.Printf("status code: %d", resp.StatusCode)
	log.Println("header: ", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("body: ", string(body))
}