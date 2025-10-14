package main

import (
	"log"
	"net/http"
	"net/rpc"
	"rpc-gob/service"
	"sync"
)

func main() {
	server := rpc.NewServer()
	server.Register(new(service.Arith))

	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	var wg sync.WaitGroup

	wg.Go(func() {
		log.Println("Starting HTTP RPC server on :1234")
		log.Fatal(http.ListenAndServe(":1234", server))
	})

	wg.Go(func() {
		log.Println("Starting HTTP RPC debug server on :1235")
		http.ListenAndServe(":1235", nil)
	})

	wg.Wait()
}
