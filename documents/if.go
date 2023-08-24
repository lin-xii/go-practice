package main

import "fmt"

func If() {
	if true {
		fmt.Println("true")
	}

	// 可以直接用表大师，还是方便的
	if i := 10; i < 0 {
		fmt.Println("i<0")
	} else {
		fmt.Println("i>0")
	}
}
