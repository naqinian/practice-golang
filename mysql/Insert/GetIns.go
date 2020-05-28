/*
* 开发人员：tss
* 创建时间：2020/5/27 6:50 下午
* 文件名称：GetIns.go
* 开发工具：GoLand
 */

package Insert

import (
	"fmt"
	"mysql/model"
)

func GetIns(){
	fmt.Println("insert func")
	r,err := model.DB.Exec("insert into person(username, sex, email)values(?, ?, ?)","张三","男","101@qq.com")
	if err != nil{
		fmt.Println("exec failed =>",err)
		return
	}

	id,err := r.LastInsertId()
	if err != nil{
		fmt.Println("exec failed =>",err)
		return
	}
	fmt.Println("insert success =>",id)

	//for r.Next(){
	//	var s string
	//	err = r.Scan(&s)
	//	if err !=nil{
	//		fmt.Println(err)
	//	}
	//	fmt.Printf("found row containing %q", s)
	//}
	//r.Close()
}


