package main

import "fmt"

func main () {
	//使用 for 循环完成下面的案例请编写一个程序，可以接收一个整数，表示层数，打印出金字塔
	//
	//编程思路
	//1 打印一个矩形
	/*

	     ***
	     ***
	     ***
	 */

	//i表示每层打印多少*
	//for j :=1; j <= 3; j++ {
	//	//j 代表每层打印多少*
	//	for j :=1; j <= 3; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//2打印半个金字塔
	/*
	*
	**
	***
	 */
	//i 表示层数
	//for i := 1; i <= 3; i++{
	//	//j表示每层打印多少
	//	for j := 1; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//3，打印整个金字
	/*
	  *
	 ***
	*****
	 */

	//i 表示层数
	//for i := 1; i <= 3; i++{
	//	//在打印*前先打印空格
	//	for k := 1; k <= 3 - i; k++{
	//		fmt.Print(" ")
	//	}
	//	//j 表示每层打印多少*
	//	for j :=1; j <= 2 * i - 1; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println("*")
	//}

	//4 将层数做成一个变量
	//var totailevei int
	//5打印空心金字塔
	/*
	  *
	 * *
	*   *
	分析，在我们给每行打印*号时，需要考虑是打印*还是打印空格
	我们的分析的结果是，每层的第一个和最后一个打印*，其他就是空的，即输入空格
	我们还分析到一个例外情况，最后层(底层） 是全部打*
	 */

	var totailevei int = 20

	//i表示层数

	for i := 1; i <= totailevei; i++ {
		//在打印*前想打印空格
		for k := 1; k <= totailevei - i; k++{
			fmt.Print(" ")
	}
	//j 表示每层打印多少*
	for j :=1; j <= 2 * i - 1; j++ {
		if j == 1 || j == 2 * i - 1 || i == totailevei{
			fmt.Print("*")
		}else {
			fmt.Print(" ")
		}
	}
	fmt.Println()}