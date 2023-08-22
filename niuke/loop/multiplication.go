package loop

import "fmt"

func multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j)
		}
		fmt.Println()
	}
}

func ExecMultiplication() {
	multiplication()
}

// 描述

// 打印9*9乘法口诀表。

// 知识点：

// golang中 for 循环中嵌套一个或多个 for 循环，代码格式如下：
// for [condition |  ( init; condition; increment ) | Range]
// {
//    for [condition |  ( init; condition; increment ) | Range]
//    {
//       statement(s)
//    }
//    statement(s)
// }
//     init： 一般为赋值表达式，给控制变量赋初值；
//     condition： 关系表达式或逻辑表达式，循环控制条件；
//     increment： 一般为赋值表达式，给控制变量增量或减量。
//     statement：循环语句

// goalng中，fmt.printf 格式化打印  %d表示数字，-3d 表示左对齐，占 3 位 \n表示下一行
