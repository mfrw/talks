package main

import "testing"
import "math/rand"

const MAX = 1000

var a []int

func init() {
	a = rand.Perm(MAX)
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Selection(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Insertion(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
	}
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})

	}
}
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Merge(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})

	}
}
func BenchmarkShell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shell(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})

	}
}
func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Qsort(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})

	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
	}
}

func BenchmarkBadSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BadSort(a)
		rand.Shuffle(MAX, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
	}
}
