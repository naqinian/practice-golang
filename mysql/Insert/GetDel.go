/*
* 开发人员：tss
* 创建时间：2020/5/27 9:55 下午
* 文件名称：GetDel.go
* 开发工具：GoLand
 */

package Insert

import (
	"fmt"
	"mysql/model"
)

func GetDel(){
	res,err := model.DB.Exec("delete from person where user_id = ?",2)
	if err != nil{
		fmt.Println("exec delete => ",err)
		return
	}

	r,err := res.RowsAffected()
	if err != nil{
		fmt.Println("rows failed => ", err)
		return
	}
	fmt.Println("delete success => ",r)

}


