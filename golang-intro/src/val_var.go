package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println(true && false)

	var a string = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "short cut for variable declaration" // HL
	fmt.Println(f)

	const n = 50000
	const l = 3e20 / n

	fmt.Println(l)
}
