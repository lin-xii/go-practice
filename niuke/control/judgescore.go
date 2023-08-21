package control

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param score int整型 分数
 * @return string字符串
 */
func judgeScore(score int) string {
	// write code here
	fmt.Println()
	switch {
	case score < 60:
		return "不及格"
	case score < 80:
		return "中等"
	case score < 90:
		return "良好"
	default:
		return "优秀"
	}
}

func ExecJudgeScore() {
	fmt.Println(judgeScore(80))
}

// 描述

// 根据成绩分数输出成绩等级，判定规则如下 分数低于60 算不及格，60-80 含60 为中等， 80-90含80  为良好，90分以上含90 为优秀。

// 知识点：

// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。 Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止。
// switch var1 {
//     case val1:
//         ...
//     case val2:
//         ...
//     default:
//         ...
// }
// 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。 您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
