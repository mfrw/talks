package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // HL
	"net/url"
)

func main() {
	res, _ := http.Get("http://www.google.com/") // HL
	if res.Request.URL.Hostname() == "auth.iiitd.edu.in" {
		magic := res.Request.URL.RawQuery
		u := res.Request.URL.String()
		res, _ = http.PostForm(u, url.Values{ // HL
			"magic":    {magic},
			"username": {"falak16018"},
			"password": {"*********"},
		})
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("Logout:", string(body[4816:4870])) // HL
	}
}
