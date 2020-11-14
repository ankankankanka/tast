package main

import (
	"fmt"
	"math"
)

func main () {
	//岳小鹏参加goland考试，他和父亲岳不群达成承诺
	//如果
	//成绩为100分时，奖励一辆车
	//成绩为（80到99）时奖励一个手机
	//成绩为（60到80）时奖励一个玩具
	//其他时，什么奖励都没有
	//请从键盘输入岳小鹏的期末考试成绩，并加以判断
	//
	//分析思路
	//1 score分数变量int
	//2 选择多分支流程控制
	//3 成绩从键盘输入，fmt.Scanln()

	var score int
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)

	//多分支判断
	//if score  == 100 {
	//	fmt.Println("奖励一辆车")
	//}else if score > 80 && score <= 99{
	//	fmt.Println("奖励一个手机")
	//}else if score >=60 && score <=80 {
	//	fmt.Println("奖励一个玩具")
	//}else {
	//	fmt.Println("什么奖励都没有")
	//}
	//使用陷阱   只会输出ok1

	//var n int = 10
	//if n > 9 {
	//	fmt.Println("ok1")//ok1输出了下面统统不执行
	//}else if n > 6{
	//	fmt.Println("ok2")
	//}else if n > 3 {
	//	fmt.Println("ok3")
	//}else {
	//	fmt.Println("ok4")
	//}

	//var b bool = true
	//if b == false {//如果写成 b = false。能编译通过么，如果能，结果是
	//	fmt.Println("a")
	//}else if b {
	//	fmt.Println("b")
	//}else if !b {
	//	fmt.Println("c")
	//}else {
	//	fmt.Println("d")
	//}

	//求ax2+bx+c=0方程的根，a,b,c分别为函数的参数，如果 b2-4ac>0,则有两个解
	//be-4ac=0,则有一个解，b2-4ac<0,则无解
	//提示1，x1=(-b+sqrt(be-4ac))/2a
	//x2=(-b-sqrt(b2-4ac))/2a
	//提示2，math.aqrt(num)；可以求平方根，需要引入 math包
	//
	//分析思路
	//1，a,b,c是三个float64
	//2，使用给出的数学公式
	//3，使用到多分支
	//4，使用math.squr方法 =》手册

	//走代码
	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0

	m := b * b - 4 * a * c
	//多分支判断
	if m > 0{
		x1 :=(-b + math.Sqrt(m))/ 2 * a
		x2 :=(-b - math.Sqrt(m))/ 2 * a
		fmt.Printf("x1=%v x2=%v",x1,x2)
	}else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v",x1)
	}else {
		fmt.Println("无解...")

	}





}
