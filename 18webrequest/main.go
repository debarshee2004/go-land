package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://example.com/"

func main() {
	fmt.Println("This is a web request example!")

	response, err := http.Get(URL)
	checkErr(err)

	fmt.Printf("Response is of type %T\n", response)
	fmt.Println(response)

	databytes, err := io.ReadAll(response.Body)
	checkErr(err)

	fmt.Println(string(databytes))

	defer response.Body.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
