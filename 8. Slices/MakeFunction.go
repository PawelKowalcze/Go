package main

//// func make([]T, len, cap) []T
//mySlice := make([]int, 5, 10)

//// the capacity argument is usually omitted and defaults to the length
//mySlice := make([]int, 5)

//mySlice := []string{"I", "love", "go"}
//fmt.Println(len(mySlice)) // 3
//fmt.Println(cap(mySlice)) // 3

import "fmt"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}

// don't edit below this line

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}
