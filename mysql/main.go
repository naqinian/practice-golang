/*
* 开发人员：tss
* 创建时间：2020/5/27 6:31 下午
* 文件名称：main.go
* 开发工具：GoLand
 */

// golang操作数据库

package main

import (
	"fmt"
	"mysql/Insert"
	_ "mysql/model"
)

func main() {
	fmt.Println("开启数据连接！")
	// 添加数据
	Insert.GetIns()
	// 查询数据
	//Insert.GetSet()
	// 修改数据
	//Insert.GetUpd()
	// 删除数据
	//Insert.GetDel()
}
