package chapter4

import "fmt"

func NilOrBlankSlice() {
	// slice 's zero value is nil!
	var s0 []int
	s1 := make([]int, 0)
	// slice can only be compared to nil
	fmt.Println(s0, s1)
	// fmt打印看着一样。但实际上，nil切片，和空切片，不一样！！！
	fmt.Println(s0 == nil) //true
	fmt.Println(s1 == nil) // false
}

func BottomArray() {
	slice := []int{1, 2, 3, 4, 5}
	subSlice := slice[1:3]
	fmt.Println("slice:", slice)
	fmt.Println("subSlice:", subSlice)

	fmt.Println("cap:", cap(subSlice))
	subSlice = append(subSlice, 100, 101, 102)
	fmt.Println("slice after append:", slice)
	fmt.Println("subSlice after append:", subSlice)

}
