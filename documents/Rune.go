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

	fmt.Println("-------------------------")
	const s = "中国"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
