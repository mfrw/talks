package main

import "fmt"

// START OMIT
type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Add(args *Args, reply *int) error { // HL
	fmt.Printf("%d + %d ", args.A, args.B)
	*reply = args.A + args.B
	fmt.Printf("==> %d\n", *reply)
	return nil
}

// END OMIT

const serverAddress = "localhost"
