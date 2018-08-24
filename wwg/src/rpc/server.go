package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// START OMIT
func main() {
	arith := new(Arith)
	rpc.Register(arith) // HL
	rpc.HandleHTTP()    // HL
	l, e := net.Listen("tcp", serverAddress+":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil) // HL
}

// END OMIT
