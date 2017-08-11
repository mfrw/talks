package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("[+] Started a webserver listening on port 8080")
	http.ListenAndServer(":8080", http.FileServer(http.Dir("."))) // HL
}
