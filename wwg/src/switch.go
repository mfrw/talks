package main

import "fmt"

func main() {
	i := 2
	fmt.Println("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2: // HL
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Noting that i know of")

	}
}
