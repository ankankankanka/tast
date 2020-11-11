package main

import "fmt"

func main() {
	//要求，可以从控制台接收用户信息，【姓名，年龄，薪水，是否通过考试】
	//方式1 fmt.Scanln()
	//1先声明需要的变量
	var name string
	var age byte
	var sal float32
	var ispass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	//当程序执行到fmt. scanln(&name)，程序会停止在这里，等待用户输入，并回车

	fmt.Println("请输入年龄")
	fmt.Scanln(&name)

	fmt.Println("请输入薪水")
	fmt.Scanln("&sal")

	fmt.Println("请输入是否通过考试")
	fmt.Scanln("&ispass")

	fmt.Printf("名字是%v \n 年龄是%v \n 是否通过考试%v\n",name,age,sal,ispass )

	//方式2 fmt.scanf 可以按指定的格式输入
	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试 ")
	fmt.Sprintf("%s %d %f %t",&name,&age,&sal,&ispass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n " )



}
