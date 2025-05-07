package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Example use of import packages and http response
	fmt.Println("Importing packages in Go")
	fmt.Println("Using net/http package to make a request")
	resp, err := http.Get("https://JSONPlaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response headers:", resp.Header)
	fmt.Println("Response body:", resp.Body)
}