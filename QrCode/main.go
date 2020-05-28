/*
* 开发人员：tss
* 创建时间：2020/5/7 2:41 下午
* 文件名称：main.go
* 开发工具：GoLand
 */

// 生成二维码

package main

import (
	"fmt"
	qrcodes "github.com/skip2/go-qrcode"
)

func main() {
	err := qrcodes.WriteFile("ceshi",qrcodes.Medium,256,"qr.png")
	if err != nil{
		fmt.Println("创建错误！")
	}
}
