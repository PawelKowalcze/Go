package main

import "fmt"

func sum(nums ...int) int {
	// nums is just a slice
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
	total := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
	// prints "6"

	names := []string{"bob", "sue", "alice"}
	printStrings(names...)
}
