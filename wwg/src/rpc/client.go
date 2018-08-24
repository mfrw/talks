package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// START OMIT
func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234") // HL
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Add", args, &reply) // HL
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith (syn): %d+%d=%d\n", args.A, args.B, reply)

	args = &Args{10, 7}
	addCall := client.Go("Arith.Add", args, &reply, nil) // HL
	_ = <-addCall.Done                                   // Check if we are done
	fmt.Printf("Arith (asyn): %d+%d=%d\n", args.A, args.B, reply)
}

// END OMIT
