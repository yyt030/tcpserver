package main

import (
	"flag"
	"log"

	"tcpserver/server"
)

var addr string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	flag.StringVar(&addr, "addr", "0.0.0.0:8888", "listen on addr")
	flag.Parse()
}

func main() {
	server.NewTcpServer(addr)
}
