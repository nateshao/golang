package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1)) // true

	str2 := "A\xfeC"
	fmt.Println(utf8.ValidString(str2)) // false

	str3 := "A\\xfeC"
	fmt.Println(utf8.ValidString(str3)) // true    // 把转义字符转义成字面值
}
