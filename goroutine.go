package main

import (
	"fmt"
	"time"
)

func output(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Goroutine() {
	output("direct")

	go output("goroutine1")

	go func(from string) {
		fmt.Println(from)
		// for i := 0; i < 3; i++ {
		// 	fmt.Println(from, ":", i)
		// }
	}("anymous")

	time.Sleep(time.Second)
	fmt.Println("done")
}
