package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	for {
		// 阻塞式从队列头部获取消息（超时时间设为0表示无限等待）
		result, err := rdb.BRPop(ctx, 0, "my_queue").Result()
		if err != nil {
			panic(err)
		}

		// result[0] 是队列名，result[1] 是消息内容
		fmt.Printf("Consumed message: %s\n", result[1])
	}
}
