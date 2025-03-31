package main

import (
	"fmt"
	"time"
)

//
//import "fmt"
//
//func hello() {
//	fmt.Println("hello bd")
//}
//func main() {
//	hello()
//	fmt.Println("main bd done!")
//}

func main() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new bd: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main bd: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
