package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Make request to the api
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer is added so that it closes the body at the end of the function
	// Just before the return
	defer res.Body.Close()
	// Defer might not be the best option if there's a lot of requests

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
