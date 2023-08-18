package niuke

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组 评委给出分数
 * @return int整型一维数组
 */
func minAndMax(s []int) []int {
	// write code here
	max := s[0]
	min := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	// 用sort.Ints能给int数组排序
	sort.Ints(s)
	fmt.Println(s)

	return []int{max, min}
}

func ExecMinAndMax() {
	scores := []int{10, 2, 1, 4, 8, 9}
	fmt.Println(minAndMax(scores))
}

// 描述

// 小明参加某个歌唱比赛，评委们进行打分,要求去掉最高分，和最低分，将最高分和最低分依次存入切片并返回。

// 知识点：

// goalng, int64最大值，最小值, 大小比较，多返回值

// golang中有符号的最大值为math.MaxInt64,最小值为math.MinInt64

// 切片的遍历有两种方式，for循环和for range循环
