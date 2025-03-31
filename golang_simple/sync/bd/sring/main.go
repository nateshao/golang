package main

import (
	"fmt"
	"strings"
)

/*
*
a:="baissbaidubbbaidu"

	b:="baidu"
	c:="baissbb"
*/
func main() {
	a := "baissbaidubbbaidu"
	b := "baidu"
	//c := "baissbb"
	res1 := replaceStr(a, b)
	fmt.Println("方法一：", res1)
	fmt.Println("-------  ---------")
	res2 := replaceStr2(a, b)
	fmt.Println("方法二：", res2)
}

func replaceStr(a string, b string) string {
	res := ""
	n := len(a)
	m := len(b)
	for i := 0; i < n; {
		// 检查当前位置是否匹配子字符串 b
		if i+m <= n && a[i:i+m] == b {
			// 如果匹配，跳过子字符串 b
			i += m
		} else {
			res += string(a[i])
			i++
		}
	}
	return res
}

func replaceStr2(a string, b string) string {
	res := strings.Replace(a, b, "", -1)
	return res
}
