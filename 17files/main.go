package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the Go programming files!")
	content := "This is a simple Go program that demonstrates how to work with files in Go."

	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes to file\n", length)

	readFiles()
	defer file.Close()
}

func readFiles() {
	databytes, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File content:")
	fmt.Println(string(databytes))
}
