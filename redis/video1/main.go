package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		//Addr: "localhost:6379",
		Addr:     "10.0.0.7:6379",
		Password: "",
		DB:       0,
	})

	arr, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接失败！", err)
		return
	}
	fmt.Println("连接成功！", arr)

	err = rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	err = rdb.Publish("rabbit", "data").Err()
	if err != nil {
		//return errors.New("发布失败")
		fmt.Println("发布失败：", err)
	}

}

func Subscribe() {
	////参数1 频道名 字符串类型
	// pubsub := redis.Redis.Subscribe("homework")
	// _, err := pubsub.Receive()
	// if err != nil {

	// 	return
	// }
	// ch := pubsub.Channel()
	// for msg := range ch {
	// 	fmt.Println(msg.Channel, msg.Payload, "\r\n")
	// }
}
