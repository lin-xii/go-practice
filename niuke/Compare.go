package niuke

import (
	"fmt"
)

func Compare(x int, y int, z int) int {
	// write code here
	prices := []int{x, y, z}
	// max := prices[0]
	// min := prices[0]
	max, min := prices[0], prices[0]
	// for i := 0; i < len(prices)-1; i++ {
	// 	if prices[i]>prices[i+1] {
	// 		max = prices[i]
	// 	}
	// }
	for _, v := range prices {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println("min:", min)
	fmt.Println("max:", max)
	return max - min
}

func ExecCompare() {
	Compare(5, 2, 3)
}
