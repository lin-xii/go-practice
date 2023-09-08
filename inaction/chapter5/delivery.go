package chapter5

import "fmt"

func ValueDelivery() {
	// value、pointer两种类型传递，应对的应该是value type
	// 而reference type，本身已经是封装的reference了。。应该不用pointer形式的传递方式
	s := []int{1, 2, 3}
	// changeLinkedItemByPointer(&s)
	changeLinkedItemByValue(s)
	fmt.Println("s:", s)
}

func changeLinkedItemByPointer(items *[]int) {
	if len(*items) > 0 {
		(*items)[0] = 1000
	}
	// fmt.Println(items[0])
	fmt.Println("func by pointer", (*items)[0])
}

func changeLinkedItemByValue(items []int) {
	if len(items) > 0 {
		items[0] = 2000
	}
	fmt.Println("func by value", items)
}
