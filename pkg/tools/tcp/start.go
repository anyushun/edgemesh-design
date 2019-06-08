package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type ProxyConfig struct {
	Addr string
	To   string
}

var conf = &ProxyConfig{}

func handleConn(conn net.Conn) {
	log.Println(conn.RemoteAddr())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("[%s]", string(netData))
	}
	resp := []byte(`HTTP/1.1 200 OK
Content-Length: 20
Content-Type: text/plain; charset=utf-8

{"name": "anyushun"}`)
	conn.Write(resp)
	conn.Close()
}

func main() {
	flag.StringVar(&conf.Addr, "addr", ":8443", "tcp server address")
	flag.Parse()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	log.Println("Server started. Press Ctrl-C to stop server")

	listen, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		log.Panicln(err)
	}

	for {
		select {
		case sig := <-ch:
			log.Println("server stop: ", sig)
		default:
			c, err := listen.Accept()
			if err != nil {
				log.Panicln(err)
			}
			handleConn(c)
		}
	}
}
