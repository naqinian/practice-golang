package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init(){
	pool = &redis.Pool{
		MaxIdle:         8,			// 最大空闲连接
		MaxActive:       0,			// 表示和数据库的最大连接。0：表示没有限制
		IdleTimeout:     100,		// 最大空闲时间
		Dial: func() (conn redis.Conn, err error) {		//初始化连接
			return redis.Dial("tcp","10.0.0.2:6379")
		},
	}
}

func main() {
	conm := pool.Get()
	defer conm.Close()

	arr,err := conm.Do("Set","name","Tom~~")
	if err != nil{
		fmt.Println("conm.Do err => :",err)
	}
	fmt.Println(arr)

	r,err := redis.String(conm.Do("Get","name"))

	fmt.Println("name => :",r)
}
