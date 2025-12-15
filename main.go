package main

import "fmt"

func main() {
	fmt.Println("Starting Go application")

	SayHello()
	result := Add(10, 20)

	fmt.Println("Addition result:", result)
}

