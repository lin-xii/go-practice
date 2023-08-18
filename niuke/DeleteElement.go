package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组 身高
 * @param index int整型 出队索引
 * @return int整型一维数组
 */
func deleteElement(s []int, index int) []int {
	// write code here

	// 如果slice底层Array非常大，要避免直接slice，使用copy更优
	// result := make([]int, 0)
	// for i, v := range s {
	// 	if i == index {
	// 		continue
	// 	}
	// 	result = append(result, v)
	// }
	// return result
	s1 := make([]int, index)
	s2 := make([]int, len(s)-index-1)

	fmt.Println(s)
	fmt.Println("start:end 3:5 左闭，右开。[ )", s[3:5])
	// 索引从0开始。所以index:3，其实是第四个元素
	fmt.Println("start: 3: []", s[3:])
	fmt.Println(":end :5 [)", s[:5])
	copy(s1, s[:index])
	copy(s2, s[index+1:])
	return append(s1, s2...)
}

func ExecDeleteElement() {
	temp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(deleteElement(temp, 2))
}

// 描述

// 学生们都排成了一队，有一个切片表示相应学生们的身高，现随机喊某个位置的人出队，返回出队后的这个切片。比如[2,3,4,5],索引为1的位置的出队，出队后切片为[2,4,5]

// 知识点：
// s|n|    切片s中索引位置为n的项
// s|:|    从切片s的索引位置0到len(s)-1 处所获得的切片
// s|low:|   从切片s的索引位置 low 到len(s)-1 处所获得的切片
// s|:high| 从切片s的索引位置 0到high 处所获得的切片，len=high
// s|low: high|  从切片s的素引位置 Iow 到high 处所获得的切片，len-high-low
// s|low: high:max|  从切片s的素引位置 low 到high 处所获得的切片，len-high-low, cap=max-low
