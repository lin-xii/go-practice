package main

import "fmt"

func For() {
	// classic
	for i := 0; i < 10; i++ {
		if i < 5 {
			// continue
			continue
		} else {
			fmt.Println(i)
		}
	}

	// most. single condition
	i := 1
	for i < 10 {
		// fmt.Println(i)
		i++
	}

	// while
	for {
		fmt.Println("while")
		break
	}
}
