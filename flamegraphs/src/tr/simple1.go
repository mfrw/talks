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

var (
	tr     = flag.String("trace", "", "Filename for tracefile")
	nrWork = flag.Int("nrwork", 45, "Max amount of work")
)

func worker(n int) {
	fib(n)
	wg.Done()
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func trackTime(s time.Time, msg string) {
	e := time.Since(s)
	fmt.Println(msg, ":", e)
}

var wg sync.WaitGroup

func main() {
	flag.Parse()
	defer trackTime(time.Now(), "MAIN")

	if *tr != "" {
		f, err := os.Create(*tr)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		trace.Start(f)
		defer trace.Stop()
	}

	wg.Add(*nrWork)
	for i := 0; i < *nrWork; i++ {
		go worker(*nrWork - i)
	}
	wg.Wait()
}
