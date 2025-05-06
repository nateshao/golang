package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 密码
		DB:       0,                // 数据库编号
	})

	ctx := context.Background()

	// 向队列尾部插入消息
	err := rdb.LPush(ctx, "my_queue", "message1").Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Message produced!")
}
