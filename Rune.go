package main

import (
	"fmt"
	"unicode/utf8"
)

func Rune() {
	ascii := "ascii"
	fmt.Println("len(ascii):", len(ascii))

	gb := "中国"
	fmt.Println("len(gb)", len(gb))
	for i, v := range gb {
		fmt.Println(i, v)
	}
	for i := 0; i < len(gb); i++ {
		// fmt.Println(gb[i])
		fmt.Printf("%x", gb[i])
		fmt.Println()
	}

	fmt.Println("中文长度：", utf8.RuneCountInString(gb))
}
