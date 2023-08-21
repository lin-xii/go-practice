package niuke

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums1 int整型一维数组
 * @param m int整型
 * @param nums2 int整型一维数组
 * @param n int整型
 * @return int整型一维数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	for i := 0; i < n; i++ {
		nums1[n+i] = nums2[i]
	}
	sort.Ints(nums1)
	return nums1
}

func ExecMerge() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{4, 5, 6}
	fmt.Println(merge(n1, 6, n2, 3))
}

// 描述

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

// 知识点：
//     1，for循环
//     2，break中断循环
//     3，append切片的追加
