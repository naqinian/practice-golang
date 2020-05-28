/*
* 开发人员：tss
* 创建时间：2020/5/27 9:31 下午
* 文件名称：GetSet.go
* 开发工具：GoLand
 */

package Insert

import (
	"fmt"
	"mysql/model"
)

func GetSet(){
	var p []model.Person
	err := model.DB.Select(&p,"select user_id,username,sex,email from person where user_id = ?",2)
	if err != nil{
		fmt.Println("exec select => ",err)
		return
	}
	fmt.Println("select succ:", p)
}


