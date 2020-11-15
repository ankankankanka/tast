
package main

import "fmt"

func main () {
	//1 统计3个班成绩情况，每个班有5名同学
	//求出各个班的平均分和所有班级的平均分，[学生的成绩从键盘输入]
	//
	//分析实现思路
	//1统计一个班的成绩情况， 每个班有5名同学，求出该班的平均分，【学生的成绩从键盘输入】=》
	//先易后难
	//2学生数就是5个，【先死后活】
	//3，声明一个sum，统计班级的总分
	//
	//
	//走代码实现

for j := 1; j <= 3; j ++{
	sum := 0.0
	for i := 1; i <= 5; i++{
		var score float64
		fmt.Printf("请输入第%的班 第%d个学生的成绩 \n", j, i )
		fmt.Scanln(&score)
		//累计总分
		sum += score
	}
	fmt.Printf("第%d个班级的平均分是%v\n",j ,sum / 5)}