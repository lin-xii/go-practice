package main

import "fmt"

func Switch() {
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoami(1)
	whoami("string")

	// type assert .(type) 只能用在switch里
	// fmt.Println("str".(type))
}
