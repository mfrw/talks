package main

import "fmt"

func main() {
	i := 2
	fmt.Println("Write", i, " as ")
	switch i { // HL
	case 1: // HL
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default: // HL
		fmt.Println("Noting that i know of")

	}
}
