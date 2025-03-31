package main

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	mu    sync.Mutex
	items []int
}

func (s *SafeSlice) Append(item int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

func (s *SafeSlice) get(index int) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.items) {
		return 0, false
	}
	return s.items[index], true
}

func main() {
	safeSlice := &SafeSlice{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeSlice.Append(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(safeSlice.items)
}
