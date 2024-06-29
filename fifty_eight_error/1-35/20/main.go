package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	char := "é"
	str := "A"
	fmt.Println(len(char))                    // 2
	fmt.Println(len(str))                     // 1
	fmt.Println(utf8.RuneCountInString(char)) // 1
	fmt.Println("cafe\u0301")                 // café    // 法文的 cafe，实际上是两个 rune 的组合
}
