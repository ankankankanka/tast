package main

import "fmt"

func main() {
	//编写一个程序，可以输入人的年龄，如果该同志的年龄大于18岁，要对
	//自己的行为负责，否则，输出"你的年龄不大这次放过你了"

	//思路分析
	//1.年龄 ==> var age int
	//2.fmt.scaln接收
	//3.if ---else

	//代码
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18{
		fmt.Scanln("你年龄大于18")
	}else {
		fmt.Println("你的年龄不大这次放过你了")
	}//if双分支基本操作

}

