package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

const MAX_NR = 1000000

var (
	p = flag.Bool("pprof", false, "Enable Profiling")
	t = flag.Bool("trace", false, "Enable Tracing")
)

// BadSort Just badly sorts data
func BadSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				s[j], s[i] = s[i], s[j]
			}
		}

	}
}

func Bubble(a []int) {
	N := len(a)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < N-1; i++ {
			if a[i+1] < a[i] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
	}
}

// Selection does a selection sort on the slice of ints and returns it
func Selection(a []int) {
	var min int
	N := len(a)
	for i := 0; i < N-1; i++ {
		min = i
		for j := i + 1; j <= N-1; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
}

// Insertion does an Insertion sort on the slice of the ints and returns a slice
func Insertion(a []int) {
	N := len(a)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func Merge(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)
	m := len(a) / 2

	for i, x := range a {
		switch {
		case i < m:
			left = append(left, x)
		case i >= m:
			right = append(right, x)
		}
	}

	left = Merge(left)
	right = Merge(right)

	return merge(left, right)
}

func merge(left, right []int) []int {

	results := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				results = append(results, left[0])
				left = left[1:len(left)]
			} else {
				results = append(results, right[0])
				right = right[1:len(right)]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:len(left)]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:len(right)]
		}
	}

	return results
}

func Shell(array []int) {
	h := 1
	for h < len(array) {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len(array); i++ {
			for j := i; j >= h && array[j] < array[j-h]; j = j - h {
				array[j], array[j-h] = array[j-h], array[j]
			}
		}
		h = h / 3
	}
}
func Qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	Qsort(a[:left])
	Qsort(a[left+1:])

	return a
}
func maxHeapify(tosort []int, position int) {
	size := len(tosort)
	maximum := position
	leftChild := 2*position + 1
	rightChild := leftChild + 1
	if leftChild < size && tosort[leftChild] > tosort[position] {
		maximum = leftChild
	}
	if rightChild < size && tosort[rightChild] > tosort[maximum] {
		maximum = rightChild
	}

	if position != maximum {
		tosort[position], tosort[maximum] = tosort[maximum], tosort[position]
		maxHeapify(tosort, maximum) //recursive
	}
}

func buildMaxHeap(tosort []int) {

	// from http://en.wikipedia.org/wiki/Heapsort
	// iParent = floor((i-1) / 2)

	for i := (len(tosort) - 1) / 2; i >= 0; i-- {
		maxHeapify(tosort, i)
	}
}

func HeapSort(tosort []int) {
	buildMaxHeap(tosort)
	for i := len(tosort) - 1; i >= 1; i-- {
		tosort[i], tosort[0] = tosort[0], tosort[i]
		maxHeapify(tosort[:i-1], 0)
	}
}

func shuffle(s []int) []int {
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}

func main() {
	flag.Parse()
	if *p {
		log.Println("Profiling Enabled")
		pf, err := os.Create("pprof.out")
		if err != nil {
			log.Fatal("Could not create pprof file")
		}
		defer pf.Close()
		pprof.StartCPUProfile(pf)
		defer pprof.StopCPUProfile()
	}

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

	defer trackTime(time.Now(), "MAIN")
	// a := rand.Perm(MAX_NR)
	a := make([]int, 0)
	for i := 0; i < MAX_NR; i++ {
		a = append(a, i)
	}
	fmt.Println("Sorting Array")
	//fmt.Println("a:=>>", a)
	Bubble(a)
	// Qsort(a)
	//Shell(a)
	fmt.Println("Sorted Array")
	//fmt.Println("a:=>>", a)
}
