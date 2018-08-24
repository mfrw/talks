package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func worker(n int) {
	fib(n)
	wg.Done()
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	start := time.Now()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(40)
	}

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println("MAIN TOOK A TOTAL OF:", elapsed)
}
