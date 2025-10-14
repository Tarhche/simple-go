package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-gob/service"
)

func main() {
	server := rpc.NewServer()
	server.Register(new(service.Arith))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Starting TCP RPC server on :1234")
	server.Accept(listener)
}
