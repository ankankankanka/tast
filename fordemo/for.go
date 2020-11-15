package main

import "fmt"

func main () {
	//golang中，有循环控制语句来处理循环的执行某段代码的方法->for 循环
	//for循环和快速入门
	//
	//
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("你好，谷歌")
	//}

	//for循环的第二种写法
	//j := 1//循环变量初始化
	//for j <=10 {
	//
	//	fmt.Println("你好，谷歌",j)
	//	j++//循环变量迭代
	//}
	//for循环的第三种写法、这种写法通常配合break使用
	//k := 1
	//for ; ; {
	//	if k <= 10 {
	//		fmt.Println("ok",k)
	//	}else {
	//		break //break 就是跳出这个for循环
	//	}
	//}
	//k++

	//字符串遍历方式，1传统方式
	var str string = "hello,world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n",str[i] )//使用到下标。。。
	}

	//字符串遍历方式 2 -for-range
	str = "abc~ok"
	for index ,val := range str {
		fmt.Printf("index=%d,val=%c \n",index,val )
	}

}
