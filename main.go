package main

import "fmt"

// Add adds two integers and returns the result.
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hi from Jenkins Go demo. 2+2 =", Add(2, 2))
}
