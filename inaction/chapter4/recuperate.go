package chapter4

import "fmt"

func Recuperate() {
	var arr = [3]int{}
	fmt.Println("array:", arr)
	// println(123)

	var arr0 [3]int
	fmt.Println(arr0)

	// 万物皆可expression
	// var arr1 [2]int = [2]int{1, 2}
	// arr2 := [2]int{1, 2}

	var s0 []int
	fmt.Println(s0)

	var sliceByMake = make([]int, 0)
	fmt.Println("slice:", sliceByMake)

	var sliceByExpression = []int{}
	fmt.Println("slice:", sliceByExpression)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)
}
