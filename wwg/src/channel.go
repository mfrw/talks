package main

import (
	"fmt"
)

func func1(channel chan string) {
	channel <- "boom"
	fmt.Println("haha")
}

func main() {
	channel := make(chan string)
	go func1(channel)
	fmt.Println(<-channel)
}
