package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("[+] Started a webserver listening on port 8080")
	http.ListenAndServe(":8080", http.FileServer(http.Dir("."))) // HL
}
