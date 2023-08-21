package niuke

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s1 int整型一维数组
 * @param s2 int整型一维数组
 * @return bool布尔型
 */
func equal(s1 []int, s2 []int) bool {
	// write code here
	if len(s1) != len(s2) {
		return false
	} else {
		isEqual := true
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[i] {
				continue
			}
			isEqual = false
		}
		return isEqual
	}
}

func ExecEqual() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println(equal(s1, s2))
}

// 描述

// 给定两个切片，判断这两个切片中的元素是否完全一样。

// 知识点：

// len(slice)求一个切片的长度

// for循环遍历切片
