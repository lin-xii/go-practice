package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param src int整型一维数组 源切片
 * @param des int整型一维数组 目的切片
 * @return int整型一维数组
 */
func sliceCopy(src []int, des []int) []int {
	// write code here
	copy(des, src)
	fmt.Println("src:", src)
	fmt.Println("des:", des)
	return des
}

func ExecSliceCopy() {
	src := []int{1, 2, 3}
	// 要小心cap、len的初始化
	// copy到des时，如果des容量过小，是copy不过去的。。。
	des := make([]int, len(src))
	fmt.Println(sliceCopy(src, des))
}

// 描述

// 给定一个切片和另一个空切片，将第一个切片复制到第二个空切片中，并返回这个被复制的空切片。

// 知识点：
// copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
