/*
* 开发人员：tss
* 创建时间：2020/5/27 10:08 下午
* 文件名称：main.go
* 开发工具：GoLand
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func main(){
	c, err := redis.Dial("tcp", "127.0.0.1:32768")
	if err != nil{
		fmt.Println("conn redis failed =>",err)
		return
	}

	fmt.Println("redis comm success")

	defer c.Close()
}


