package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type ServerConfig struct {
	CA   string
	Cert string
	Key  string
}

var conf *ServerConfig = &ServerConfig{}

func main() {
	var addr string = ""
	flag.StringVar(&conf.CA, "ca", "", "ca for verify server")
	flag.StringVar(&conf.Cert, "cert", "", "cert for tls")
	flag.StringVar(&conf.Key, "key", "", "key for tls")
	flag.StringVar(&addr, "addr", ":8443", "server listen address")
	flag.Parse()

	tlsConfig := &tls.Config{}

	if conf.CA != "" {
		caData, err := ioutil.ReadFile(conf.CA)
		if err != nil {
			log.Panic(err)
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caData)
		tlsConfig.ClientCAs = pool
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	if conf.Cert != "" && conf.Key != "" {
		pair, err := tls.LoadX509KeyPair(conf.Cert, conf.Key)
		if err != nil {
			log.Panicln(err)
		}
		tlsConfig.Certificates = []tls.Certificate{pair}
	}

	s := http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.String())
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.Write(body)
		}),
		TLSConfig: tlsConfig,
	}
	log.Println("Server started. Press Ctrl-C to stop server")
	if conf.CA != "" {
		go s.ListenAndServeTLS("", "")
	} else {
		go s.ListenAndServe()
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	log.Println(<-ch)
	log.Println("exit")
	signal.Stop(ch)
}
