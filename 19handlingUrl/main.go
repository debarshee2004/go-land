package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://example.com:5000/learn?c=react&id=123456"

func main() {
	fmt.Println("Welcome to the handling urls section")
	fmt.Println(URL)

	// parsing
	result, _ := url.Parse(URL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
}
