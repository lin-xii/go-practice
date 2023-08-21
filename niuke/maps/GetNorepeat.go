package maps

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组
 * @return int整型一维数组
 */
func getNoRepeat(s []int) []int {
	// write code here
	norepeat := make(map[int]int)

	for _, v := range s {
		val, ok := norepeat[v]
		if ok {
			norepeat[v] = val + 1
		} else {
			norepeat[v] = 1
		}
	}

	keys := make([]int, 0, len(norepeat))
	for k, v := range norepeat {
		if v == 1 {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)
	return keys
}

func ExecGetNoRepeat() {
	params := []int{1, 2, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9}
	fmt.Println(getNoRepeat(params))
}

// 描述

// 给定一个数组，找出数组中所有不重复的数字，并按照从小到大的顺序进行输出

// 知识点：
//     1,map用make方式进行初始化
//     2,切片可以用[]int{}的方式进行初始化
//     3,for range遍历切片
//     4,_,ok :=map[key]的方式判断m中的key是否存在
//     5,切片用append方式进行追加
