package main
import"fmt"

func main () {
	//打印出九九乘法表
	//i 表示层数
	var num int = 20
	for i := 1; i <= num; i++{
		for j := 1; j <= i; j++{
			fmt.Printf("%v * %v = %v \t",j, i, j * i)
		}
		fmt.Println()
	}
}