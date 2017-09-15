package main

import "fmt"

func main() {

	t := []string{"a", "b", "c"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) // HL
	for i := 0; i < 3; i++ { // HL
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // HL
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}
