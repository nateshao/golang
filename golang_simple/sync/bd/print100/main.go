package main

import (
	"fmt"
	"sync"
)

func main() {
	// 打印1-100
	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 100; i += 2 {
			ch1 <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i <= 100; i += 2 {
			ch2 <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			if i%2 != 0 {
				fmt.Println(<-ch1)
			} else {
				fmt.Println(<-ch2)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
