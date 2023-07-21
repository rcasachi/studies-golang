package main

import (
	"fmt"
	"time"
)

func processing() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	// go processing()
	// go processing()
	// processing()

	channel := make(chan int)

	go func() {
		// channel <- 1 // T2
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	// fmt.Println(<-channel)
	// time.Sleep(time.Second * 2)

	// for x := range channel {
	// 	fmt.Println(x)
	// 	time.Sleep(time.Second)
	// }
	go worker(channel)
	go worker(channel)
	go worker(channel)
	worker(channel)
}

func worker(channel chan int) {
	for {
		fmt.Println(<-channel)
		time.Sleep(time.Second)
	}
}
