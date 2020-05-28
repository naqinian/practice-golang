/*
* 开发人员：tss
* 创建时间：2020/5/28 12:37 下午
* 文件名称：main.go
* 开发工具：GoLand
 */

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:32768")
	if err != nil{
		fmt.Println("conn redis failed =>",err)
		return
	}

	fmt.Println("redis comm success")

	defer c.Close()

	//_,err = c.Do("MSet","abc",100,"efg",200)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}

	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}
