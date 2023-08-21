package maps

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return char字符型
 */
func character(s string) byte {
	// write code here
	for i := 0; i < len(s); i++ {
		temp := s[i]
		fmt.Println(temp)
	}

	dict := make(map[rune]byte)
	for _, v := range s {
		// fmt.Println(v)
		// fmt.Println(string(v))
		val, ok := dict[v]
		if ok {
			dict[v] = val + 1
		} else {
			dict[v] = 1
		}
	}

	var mostKey rune
	for key, v := range dict {
		if v > dict[mostKey] {
			mostKey = key
		}
	}

	fmt.Println(dict)
	fmt.Println(mostKey)

	return byte(mostKey)
}

func ExecCharacter() {
	params := "hello"
	fmt.Println(character(params))
}

// 描述

// 给定一个只由字母和数字组成的字符串,，统计每个字符出现的次数，并返回出现次数最多的字符。

// 知识点：
// 1，组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
// var a = '中'
// 1，map的每个key是唯一的
// 2，map的声明：map[KeyType]ValueType KeyType:表示键的类型。ValueType:表示键对应的值的类型。map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：make(map[KeyType]ValueType, [cap]) 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
