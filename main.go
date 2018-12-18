package main

import (
	"flag"

	"github.com/edouardparis/spark/server"
)

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "host", ":8080", "server listen address")
	flag.Parse()

	server.New(listenAddr).Run()
}
