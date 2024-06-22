package main

import (
	"fmt"
)

const (
	beijing = iota
	shanghao
	shenzhen
)

// 全局变量m
var m = 100

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
	fmt.Println(beijing)
	fmt.Println(shanghao)
	fmt.Println(shenzhen)
}
