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
	stream := "my_stream"
	group := "my_group"
	consumer := "consumer1"

	// 创建消费者组（如果不存在）
	err := rdb.XGroupCreateMkStream(ctx, stream, group, "0").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		panic(err)
	}

	for {
		// 从消费者组读取消息（阻塞式）
		entries, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  []string{stream, ">"}, // ">" 表示读取未分配给其他消费者的消息
			Count:    1,
			Block:    0, // 无限阻塞
		}).Result()

		if err != nil {
			panic(err)
		}

		for _, entry := range entries[0].Messages {
			// 处理消息
			fmt.Printf("Consumed message: %v\n", entry.Values["data"])

			// 确认消息已处理（ACK）
			err = rdb.XAck(ctx, stream, group, entry.ID).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}
