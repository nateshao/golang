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

	// 添加消息到 Stream
	err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "my_stream",
		Values: map[string]interface{}{"data": "message1"},
	}).Err()

	if err != nil {
		panic(err)
	}
	fmt.Println("Message produced to stream!")
}
