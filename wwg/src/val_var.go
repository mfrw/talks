package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println(true && false)

	var a string = "initial" // HL
	fmt.Println(a)
	var b, c int = 1, 2 // HL
	fmt.Println(b, c)
	var d = true // HL
	fmt.Println(d)
	var e int // HL
	fmt.Println(e)
	f := "short variable declaration" // HL
	fmt.Println(f)

	const n = 50000
	const l = 3e20 / n
	fmt.Println(l)
}
