package control

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param forwards string字符串 推箱子方向
 * @return bool布尔型
 */
func pushBox(forwards string) bool {
	// write code here
	x, y := 0, 0
	for _, v := range forwards {
		switch string(v) {
		case "R":
			x += 1
		case "L":
			x -= 1
		case "U":
			y += 1
		case "D":
			y -= 1
		}
	}
	fmt.Println(x, y)
	return x == 0 && y == 0
}

func ExecPushbox() {
	fmt.Println(pushBox("UDL"))
}

// 描述

// 在二维平面上，有一个人在推箱子，假设从原点 (0, 0) 开始。给出它的移动顺序，判断这个箱子在完成移动后是否在 (0, 0) 处结束。移动顺序由字符串 moves 表示。字符 move[i] 表示其第 i 次移动。推箱子的有效动作有 R（右），L（左），U（上）和 D（下）。如果推箱子完成所有动作后返回原点，则返回 true。否则，返回 false。
