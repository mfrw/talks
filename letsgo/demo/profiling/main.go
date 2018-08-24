package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"sort"
	"time"
)

const size = 1000

// shuffle the slice
func shuffle(s []int) []int {
	for i := len(s) - 1; i > 0; i-- {
		//j := rand.Int63() % int64(i+1)
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// using append
func initSlice(s []int, n int) []int {
	for i := 1; i < n+1; i++ {
		s = append(s, i)
	}
	return s
}

// Algorithms police: just passing pointers around
func badSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				s[j], s[i] = s[i], s[j]
			}
		}
		//		fmt.Println(s)
	}
	return s
}

func goodSort(s []int) []int {
	sort.Ints(s)
	return s
}

func timeTrack(start time.Time, msg string) {
	elapsed := time.Since(start)
	fmt.Println(msg, ":", elapsed)
}

func main() {
	defer timeTrack(time.Now(), "MAIN")

	pprof.StartCPUProfile(os.Stderr)
	defer pprof.StopCPUProfile()
	//trace.Start(os.Stderr)
	//defer trace.Stop()
	var slice []int
	slice = make([]int, 0)
	fmt.Println(slice)
	fmt.Println("")

	/* intit slice */
	slice = initSlice(slice, size)
	fmt.Println(slice)
	fmt.Println("")

	shuffle(slice)
	fmt.Println(slice)
	fmt.Println("")

	/* insane loop */
	for i := 0; i < 2000; i++ {
		shuffle(slice)
		//goodSort(slice)
		badSort(slice)
	}

	fmt.Println(slice)
	fmt.Println("")
}
