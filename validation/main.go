// 各种验证

package main

import "regexp"

func main()  {
	//验证身份证
	getCard()
	//验证手机号
	getPhone()
	//验证邮箱
	getEmail()
	//验证中文
	getChinese()
}

func getCard()  {
	usercard15 := `123456789012345`
	usercard18 := `12345678901234567x`

	// 15位身份证号码：15位全是数字
	if isorno, _ := regexp.MatchString(`^(\d{15})$`, usercard15); isorno {
		println(`正确的15位身份证`)
	} else {
		println(`错误的15位身份证`)
	}

	// 18位身份证：前17位为数字，第18位为校验位，可能是数字或X
	if isorno, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, usercard18); isorno {
		println("正确的18位身份证")
	} else {
		println("错误的18位身份证")
	}
}

func getPhone()  {
	mobile := `13200000000`

	// 识别手机号码
	isorno, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if isorno {
		println(`正确手机号`)
	} else {
		println(`错误的手机号`)
	}
}

func getEmail(){
	email := `heydon-lee@qq.com`

	// 识别电子邮件地址
	isorno, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	if isorno {
		println(`正确的邮箱`)
	} else {
		println(`不是邮箱`)
	}
}

func getChinese()  {
	chinese := `你好啊`

	// 识别全是中文
	isorno, _ := regexp.MatchString(`[\x{4e00}-\x{9fa5}]+`, chinese)
	if isorno {
		println(`这段话含有中文`)
	} else {
		println(`这段话没有中文`)
	}
}