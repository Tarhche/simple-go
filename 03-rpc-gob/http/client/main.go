package main

import (
	"flag"
	"log"
	"net/rpc"
	"rpc-gob/service"
	"strconv"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 1234, "port to connect to")
	flag.Parse()

	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()

	syncCall(client)
	asyncCall(client)
}

func syncCall(client *rpc.Client) {
	args := &service.Args{
		A: 7,
		B: 8,
	}

	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

func asyncCall(client *rpc.Client) {
	quotient := new(service.Quotient)
	args := &service.Args{
		A: 7,
		B: 8,
	}

	divCall := client.Go("Arith.Divide", args, quotient, nil)

	<-divCall.Done // will be equal to divCall

	if divCall.Error != nil {
		log.Fatal("arith error:", divCall.Error)
	}
	log.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}
