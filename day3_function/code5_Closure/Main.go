package main

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-1)
}

/**
递归，闭包
*/
func main() {
	c := a()
	c()
	c()
	c()

	a() //不会输出i

	// 斐波那契数列(Fibonacci)
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fibonaci(i))
	}
}
