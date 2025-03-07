package main

import "fmt"

// syntax for _ in range
//for INDEX, ELEMENT := range SLICE { }

func main() {
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
	// 0 apple
	// 1 banana
	// 2 grape
}
