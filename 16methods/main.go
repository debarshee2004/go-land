package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Go programming methods!")
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetAge() int {
	return u.Age
}
