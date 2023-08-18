package niuke

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param length int整型 切片初始化长度
 * @param capacity int整型 切片初始化容量
 * @return int整型一维数组
 */
func makeslice(length int, capacity int) []int {
	// write code here
	s := make([]int, length, capacity)

	// for i, v := range s {
	// 	v = i
	// }

	for i := range s {
		// val = i
		s[i] = i
		// fmt.Println(i, val)
	}

	return s
}

func ExecMakeSlice() {
	fmt.Println(makeslice(2, 4))
}

// 描述

// 创建一个制定长度，容量的int类型切片，设置该切片的每个位置的值等于其索引值,最后返回该切片。

// 知识点；

// 切片的声明：var identifier []type
// make() 函数来创建切片:var slice1 []type = make([]type, len) 指定容量，其中 capacity 为可选参数：make([]T, length, capacity)

// for循环遍历切片
// for init; condition; post { }
//     for condition { }
//     for { }
//     init： 一般为赋值表达式，给控制变量赋初值；
//     condition： 关系表达式或逻辑表达式，循环控制条件；
//     post： 一般为赋值表达式，给控制变量增量或减量。
//     for语句执行过程如下：
//     ①先对表达式 init 赋初值；
//     ②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句

// len(slice)获取切片长度，cap(slice)获取切片容量
