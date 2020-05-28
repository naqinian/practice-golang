/*
* 开发人员：tss
* 创建时间：2020/5/27 6:59 下午
* 文件名称：suface.go
* 开发工具：GoLand
 */

package model

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}



