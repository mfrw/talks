package main

import "testing"

var tslice []int

func init() {
	tslice = make([]int, 200)
	for i := 199; i > -1; i-- {
		tslice[i] = i
	}
}

func BenchmarkGoodSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goodSort(tslice)
	}
}

func BenchmarkBadSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		badSort(tslice)
	}
}
