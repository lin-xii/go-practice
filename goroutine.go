package main

import (
	"fmt"
	// "time"
)

func output(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		// msg <- "ping"
	}
}

func Goroutine() {
	// msg := make(chan string)
	msg := make(chan []string)
	output("direct")

	go output("goroutine1")

	go func(from string) {
		fmt.Println(from)
		// for i := 0; i < 3; i++ {
		// 	fmt.Println(from, ":", i)
		// }
		// msg <- "ping"
		// msg <- append(<-msg, "str")
		// chan可以是slice。但是只能存一次，不能append
		msg <- []string{"1", "2"}
	}("anymous")

	// time.Sleep(time.Second)
	fmt.Println("done")
	fmt.Println(<-msg)
}
