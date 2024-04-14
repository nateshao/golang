package main

import (
	"fmt"
	"time"
)

//
//import "fmt"
//
//func hello() {
//	fmt.Println("hello goroutine")
//}
//func main() {
//	hello()
//	fmt.Println("main goroutine done!")
//}

func main() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
