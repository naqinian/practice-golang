package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "10.0.0.7:6379")
	if err != nil {
		fmt.Println("con redis falled", err)
		return
	}
	fmt.Println("redis con success")
	defer c.Close()

	// _, err = c.Do("Set", "rabbit", 100)
	// if err != nil {
	// 	fmt.Println("Set rabbit err：", err)
	// 	return
	// }

	f, err := redis.Int(c.Do("Get", "rabbit"))
	if err != nil {
		fmt.Println("Get rabbit err", err)
	}
	fmt.Println(f)
}
