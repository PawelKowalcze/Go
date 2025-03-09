package main

import (
	"fmt"
	"time"
)

//for item := range ch {
//// item is the next value received from the channel
//}

func concurrrentFib(n int) {
	chInts := make(chan int)
	go func() {
		fibonacci(n, chInts)
	}()
	for v := range chInts {
		fmt.Println(v)
	}
}

// TEST SUITE - Don't touch below this line

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
