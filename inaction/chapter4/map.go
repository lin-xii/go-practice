package chapter4

import "fmt"

func SomeMap() {
	var m map[string]int
	// fmt.Println(m, m == nil)

	m = map[string]int{"one": 1}

	fmt.Println(m)
}
