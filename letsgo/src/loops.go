package main

import "fmt"

func main() {
	i := 0
	for i <= 3 { // HL
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ { // HL
		fmt.Println(j)
	}
	// var j cant be used outside the loop body
	// its scope is limited to the loop only

	for { // HL
		fmt.Println("Infinite loop")
		break
	}
}
