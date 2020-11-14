package main

import "fmt"

func main() {
	//编写程序，声明2个int32形变量并赋值，判断两数之和，如果大于等于50，打印'hello world'

//	分析
//	1.变量
//	2.单分支

	var n1 int32 =10
	var n2 int32 = 50
	if n1 + n2 > 50{
		fmt.Println("hello world")
	}

	//编写程序，声明2个float64型变量并赋值，判断第一个数大于10.0
	//且第2个数小于20.0 打印俩数之和

	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("和=",(n3 + n4))

	}

}