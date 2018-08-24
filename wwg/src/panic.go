package main

import "fmt"

// START OMIT
func main() {
	f()
	fmt.Println("Returned normally from f.")
}
func f() {
	defer func() { // HL
		if r := recover(); r != nil { // HL
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // NEVER PRINTED
}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i)) // HL
	}
	defer fmt.Println("Defer in g", i) // HL
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// END OMIT
