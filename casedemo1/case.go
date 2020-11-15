package main
import (
	"fmt"

)

func main () {
	//1使用switch 把小写类型的 char型转为大写（键盘输入）
	//只转换 a,b,c,d,e,其他的输出"other"
	//var char byte
	//fmt.Println("请输入一个字符。。")
	//fmt.Scanf("%c", &char)
	//
	//switch char {
	//case'a' :
	//	fmt.Println("A")
	//case 'b' :
	//	fmt.Println("B")
	//case 'c' :
	//	fmt.Println("C")
	//case 'E' :
	//	fmt.Println('E')
	//default :
	//	fmt.Println("other")
	//}

	//var score float64
	//fmt.Println("请输入成绩")
	//fmt.Scanln(&score)
	//
	//switch int  (score / 60){
	//case 1 :
	//	fmt.Println("及格")
	//case 0 :
	//	fmt.Println("不及格")
	//default :
	//	fmt.Println("输入有误")
	//}

	//3 根据用户指定月份
	//打印该月份所属的季节，3，4，5，春季 6，7，8 夏季9，10，11秋季12， 1，2，冬季

	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5 :
		fmt.Println("spring")
	case 6, 7, 8 :
		fmt.Println("summer")
	case 9, 10,11 :
		fmt.Println("autumn")
	case 12, 1, 2 :
		fmt.Println("winter")
	default :
		fmt.Println("输入有误")
	}
	}
