package keyvalue

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Run keyvalue
func Run(client *redis.Client) {
	var val interface{}
	var err error
	//  Set 第三个参数代表key的过期时间，0代表不会过期。
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err = client.Get("key").Result()
	// 判断查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("Get-->key", val)

	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err = client.GetSet("key", "new value").Result()

	if err != nil {
		panic(err)
	}
	// 打印key的旧值
	fmt.Println("GetSet设置key一个新值，并返回一个旧值-->", val)

	// 第三个参数代表key的过期时间，0代表不会过期。
	err = client.SetNX("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("SetNX如果key不存在，则设置这个key的值")

	// 批量设置key的值
	err = client.MSet("key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("MSet批量设置key的值")

	// MGet函数可以传入任意个key，一次性返回多个值。
	// 这里Result返回两个值，第一个值是一个数组，第二个值是错误信息
	val, err = client.MGet("key1", "key2", "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("MGet批量取值-->", val)

	// 删除key
	client.Del("key")

	// 删除多个key, Del函数支持删除多个key
	err = client.Del("key1", "key2", "key3").Err()
	if err != nil {
		panic(err)
	}

	// 设置key的过期时间,单位秒
	client.Expire("key", 3)
}
