package niuke

import "fmt"

func Array() []int {
	score := [5]int{2, 5, 4, 6, 5}
	return score[:]
}

func ExecArray() {
	fmt.Println(Array())
}

// 描述

// 小明投了5次保龄球，每次的分数分别为2，5，4，6，5,用一个数组记录这5次分数，然后输出这个数组。

// 知识点 ：

// 数组：是同一种数据类型的固定长度的序列。

// 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。

// 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
