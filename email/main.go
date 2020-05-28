/*
* 开发人员：tss
* 创建时间：2020/5/19 6:16 下午
* 文件名称：main.go
* 开发工具：GoLand
 */


// 发送邮件

package main

import (
	"github.com/go-gomail/gomail"
)


func main(){
	useradd := "123123@qq.com"
	username := "XXX网站邮件通知"
	name := "123123qq.com"
	pass := "123123"

	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", useradd, username)
	// 收件人
	m.SetHeader("To",
		m.FormatAddress("123123@qq.com", "123"),
	)
	// 主题
	m.SetHeader("Subject", username)
	// 正文
	m.SetBody("text/html", "123123")
	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer("smtp.qq.com", 465, name, pass)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}


