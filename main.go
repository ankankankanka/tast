
package main

import "fmt"

func main() {
	//案例
	//请编写一个程序，该程序可以接收一个字符，比如，a,b,c,d,e,f,g,
	//a表示星期一，b表示星期二，根据用户的输入显示相依的信息
	//
	//要求使用；switch 语句完成
	//
	//分析思路
	//1，定义一个变量接收字符
	//2，使用switch完成
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g,")
	fmt.Scanf("%c",&key)

	switch key {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子爬雪山")
	default :
		fmt.Println("输入有误...")

	}
}
