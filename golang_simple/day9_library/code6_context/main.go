package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//func worker() {
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//	}
//	// 如何接收外部命令实现退出
//	wg.Done()
//}
//
//func main() {
//	wg.Add(1)
//	go worker()
//	wg.Wait()
//	fmt.Println("over")
//}

// 全局变量
var exit bool

// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}
func main() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exit = true                 // 修改全局变量实现子goroutine的退出
	wg.Wait()
	fmt.Println("over")
}
