package maps

import (
	"fmt"
	"unicode/utf8"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param des string字符串
 * @param src string字符串
 * @return bool布尔型
 */
func canConstruct(des string, src string) bool {
	// write code here
	// 这个适用于src字符无限次使用。。题目里要求只能用一次。。。
	// result := true
	// for _, v := range des {
	// 	if strings.ContainsRune(src, v) {
	// 		continue
	// 	}
	// 	result = false
	// 	break
	// }
	// return result

	temp := make(map[rune]rune)
	sliceSrc := []rune(src)
	// 可以直接遍历des，而不用特意转换成sliceSrc
	for _, v := range des {
		for _, vSrc := range sliceSrc {
			if v == vSrc {
				temp[vSrc] = vSrc
			}
		}
	}
	fmt.Println(temp)
	fmt.Println(len(temp))

	return utf8.RuneCountInString(des) == len(temp)
}

func ExecCanConstruct() {
	des := "aa"
	src := "ab"
	fmt.Println(canConstruct(des, src))
}

// 描述

// 给定两个字符串des 和src  ，判断 des能不能由 src 里面的字符构成,//如果可以，返回 true ；否则返回 false,src中的每个字符只能在 des 中使用一次。

// 知识点：
// for range遍历字符串
// cnt[ch-'a']隐士转换byte转为int类型
// 数组的索引可以充当一个map的key，来表示唯一
