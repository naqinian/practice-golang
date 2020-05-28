/*
* 开发人员：tss
* 创建时间：2020/5/27 6:51 下午
* 文件名称：models.go
* 开发工具：GoLand
 */

package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

// 创建一个mysql连接池
var DB *sqlx.DB

func init(){
	database,err := sqlx.Open("mysql","root:12345678@tcp(127.0.0.1:3306)/text")

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)

	if err != nil{
		fmt.Println("数据库连接失败！")
		return
	}
	DB = database
	fmt.Println("数据库连接成功！")
}

