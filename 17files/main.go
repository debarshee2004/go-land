package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File Management in Go!")

	filename := "example.txt"

	// ----------------------------
	// Example 1: Creating and writing to a file
	fmt.Println("\nExample 1: Creating and writing to a file")
	file, err := os.Create(filename) // creates or truncates if file exists
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // ensure file is closed after we're done

	content := "Hello, Go File I/O!\nThis is line 2.\n"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully.")

	// ----------------------------
	// Example 2: Appending to an existing file
	fmt.Println("\nExample 2: Appending to a file")
	appendFile, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		return
	}
	defer appendFile.Close()

	_, err = appendFile.WriteString("This is an appended line.\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("File appended successfully.")

	// ----------------------------
	// Example 3: Reading entire file into memory
	fmt.Println("\nExample 3: Reading entire file")
	data, err := os.ReadFile(filename) // returns []byte
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content as string:")
	fmt.Println(string(data))

	// ----------------------------
	// Example 4: Reading line-by-line using bufio
	fmt.Println("\nExample 4: Reading line by line")
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading lines:", err)
	}

	// ----------------------------
	// Example 5: Copying file content to another file
	fmt.Println("\nExample 5: Copying content to another file")
	src, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer src.Close()

	dst, err := os.Create("copy.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dst.Close()

	bytesCopied, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Printf("Copied %d bytes to 'copy.txt'\n", bytesCopied)

	// ----------------------------
	// Example 6: Checking if a file exists
	fmt.Println("\nExample 6: Checking file existence")
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File exists.")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("Error checking file:", err)
	}

	// ----------------------------
	// Example 7: Deleting a file
	fmt.Println("\nExample 7: Deleting 'copy.txt'")
	err = os.Remove("copy.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("'copy.txt' deleted successfully.")
	}
}
