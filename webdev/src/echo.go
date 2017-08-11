package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // HL
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler) // HL
	log.Println("[+] Server listening on 8080")
	http.ListenAndServe(":8080", nil) // HL
}
