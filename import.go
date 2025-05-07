package main

import "fmt"
import foo "net/http"

func main() {
	fmt.Println("Hello, Go Standard Library!")
	fmt.Println("This is a simple HTTP server.")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
}
