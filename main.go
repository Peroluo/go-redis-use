package main

import (
	"go-redis-use/keyvalue"

	"github.com/go-redis/redis"
)

var client *redis.Client

// 1.redis连接
func createClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
}

func main() {
	// 连接redis
	createClient()
	// 2.redis keyvalue基本方法
	keyvalue.Run(client)
}
