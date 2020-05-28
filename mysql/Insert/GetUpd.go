/*
* 开发人员：tss
* 创建时间：2020/5/27 9:44 下午
* 文件名称：GetUpd.go
* 开发工具：GoLand
 */

package Insert

import (
	"fmt"
	"mysql/model"
)

func GetUpd(){
	res,err := model.DB.Exec("update person set username=? where user_id=?", "stu0003", 2)
	if err != nil{
		fmt.Println("exec update => ",err)
		return
	}

	msg,err := res.RowsAffected()
	if err != nil{
		fmt.Println("rows failed => ", err)
		return
	}
	fmt.Println("update success => ",msg)
}


