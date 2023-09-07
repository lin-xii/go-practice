package chapter4

import "fmt"

func SomeArray() {
	arr := [3]int{1, 2, 3}
	arr[1] = 4
	fmt.Println(arr)

}
