package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	var result []int

	go func() {
		for item := range ch {
			result = append(result, item)
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	wg.Wait()
	close(ch)
	fmt.Println(result)
}
