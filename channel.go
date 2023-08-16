package main

import "fmt"

func Channel() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	fmt.Println(<-message)
}
