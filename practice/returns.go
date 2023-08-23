package practice

import "fmt"

func multiReturns() (first int, second float32) {
	defer func() {
		second = 3.0
	}()
	first = 1
	second = 2.0
	return
}

func RunMultiReturns() {
	fmt.Println(multiReturns())
}
