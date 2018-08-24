package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var wg sync.WaitGroup
var t = flag.Bool("trace", false, "Enable Tracing")

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib_worker(n int) {
	defer wg.Done()
	fib(n)
}

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}

// START OMIT
func main() {
	flag.Parse()
	if *t {
		log.Println("Tracing Enabled")
		tf, err := os.Create("trace.out")
		if err != nil {
			log.Fatal("Could not create trace file")
		}
		defer tf.Close()
		trace.Start(tf)
		defer trace.Stop()
	}

	defer trackTime(time.Now(), "MAIN") // HL
	MAX := 64
	FIB := 35
	for i := 0; i < MAX; i++ {
		wg.Add(1)
		go fib_worker(FIB) // HL
	}
	wg.Wait()
	fmt.Println("Total Times FIB Executed:", MAX)
}

// END OMIT
