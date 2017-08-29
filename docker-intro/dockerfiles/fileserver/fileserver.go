//author: mfrw
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// usage help --help or -help
func main() {
	/* A static web server serving a current path
	It can also be used a hacky file tx program.
	Just build it with with the arch you want to
	run on, a static ready to run binary is the output:

	GOOS=windows go build fileserver.go // for windows
	GOOS=openbsd go build fileserver.go // for openbsd
	GOOS=linux   go build fileserver.go // for linux
	*/

	if len(os.Args) < 2 {
		fmt.Println("Please run with --help for more options")
	}

	tpath, err := os.Getwd() // get the current working dir
	if err != nil {
		fmt.Fprintf(os.Stderr, "Some bad stuff occured\n")
		os.Exit(-1)
	}

	// get the flags from user
	path := flag.String("path", tpath, "Path of the directory")
	port := flag.String("port", "8080", "Port on which to Host")
	flag.Parse()

	fmt.Println("Server starting on Port " + *port)
	fmt.Println("Server Hosting: ", *path)
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*path))))
}
