package niuke

import "fmt"

func wardrobe() {
	cloth := []string{"帽子", "围巾", "衣服", "裤子", "袜子"}
	fmt.Println(cloth)
}

func ExecWardrobe() {
	wardrobe()
}

// 描述

// 有一个置衣柜，有5层，第一层放的都是 "帽子"  ，第二层 放的都是 "围巾" , 第三层放的 都是 "衣服" ，第四层放的都是 "裤子"  第五层放的都是 "袜子",用一个切片表示这个衣柜每层放置的物品，并输出这个切片。

// 知识点：

// 切片的声明：var identifier []type

// make() 函数来创建切片 : var slice1 []type = make([]type, len) 指定容量，其中 capacity 为可选参数：make([]T, length, capacity)

// append追加切片
// 输入描述：

// 无
// 输出描述：

// [帽子 围巾 衣服 裤子 袜子]
