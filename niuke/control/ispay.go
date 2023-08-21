package control

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param hight double浮点型 身高
 * @return bool布尔型
 */
func ispay(hight float64) bool {
	// write code here
	fmt.Print()
	return hight >= 160.0
}

func ExecIspay() {
	fmt.Println(ispay(111))
}

// 描述

// 有个游乐园，如果身高小于160.0，则可以免费入场，身高大于等于160.0岁则要收费，根据输入的身高判断是否需要收费，收费返回true,不收费返回false。

// 知识点：

// 条件语句是用来判断给定的条件是否满足(表达式值是否为0)，并根据判断的结果(真或假)决定执行的语句。

// golang中 条件语句主要指if语句 ，if 语句 由一个布尔表达式后紧跟一个或多个语句组成。if 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则不执行。格式如下：
//      if 布尔表达式 {
//     /* 在布尔表达式为 true 时执行 */
//     }

// golang中 float浮点型 可以表示小数

// golang中 关系运算符 ==,<,>,<=,>=,!=，分别表示 等于，小于，大于，小于等于，大于等于，不等于
