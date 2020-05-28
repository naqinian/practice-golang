package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

const URL = "127.0.0.1:27017"

func main() {
	session, err := mgo.Dial(URL)

	if err != nil {
		fmt.Println("数据库连接错误！")
	}

	c := session.DB("test").C("marios")

	c.Insert(map[string]interface{}{"name": "如初", "age": 25})

	var result []map[string]interface{}

	c.Find(nil).All(&result)

	fmt.Println(result)
}
