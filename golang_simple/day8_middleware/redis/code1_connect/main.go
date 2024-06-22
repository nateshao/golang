package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 链接Redis
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis conn fail", err)
		return
	}
	fmt.Println("redis conn success")
	defer c.Close()
}
