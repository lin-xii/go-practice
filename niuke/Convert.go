package niuke

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组
 * @return int整型一维数组
 */
func convert(s []int) []int {
	// write code here
	// 方案1
	// result := make([]int, len(s))
	// for i := len(s) - 1; i >= 0; i-- {
	// 	result[len(s)-i-1] = s[i]
	// }
	// return result

	// 方案2
	// 时间复杂度为原来的1/2，空间复杂度也下降了一半。确实更好一些
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s

}

func ExecConvert() {
	params := []int{1, 3, 2, 3, 4, 6}
	fmt.Println(convert(params))
}

// 描述

// 小朋友们依次站成了一排，现要将他们调换顺序，反着排，按照从最右的人站在最左边，倒数最右边的人站在最左边第二个位置，以此类推。比如小朋友的顺序为[1,3,2,3,4,6]，重新排列后为[6,4,3,2,3,1]

// 知识点：

// len(slice)求一个切片的长度

// for循环遍历切片

// golang提供了多重赋值的特性可以轻松实现变量的交换，变量一，变量二 ：= 变量二，变量一
