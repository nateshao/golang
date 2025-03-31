package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
}

var p Person
var mu sync.Mutex

func update(name string, age int) {
	mu.Lock()
	defer mu.Unlock()
	p.name = name
	time.Sleep(time.Millisecond * 200)
	p.age = age
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			update(name, age)
		}()
	}
	wg.Wait()
	fmt.Printf("p.name= %s\np.age=%v", p.name, p.age)
}
