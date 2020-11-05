package main

import "fmt"

//演示golang中基本数据类型的转换
func main () {
	var i  = 100
	//希望将i => float64
	var n1  = float32(i)
	var n2 = int8(i)
	var n3 = int64(i)//低精度->高精度

	fmt.Printf("i=%v n1=%v n2=%v n3=%v",i,n1,n2,n3)

	//在转换中，比如将 int64 转成 int8 【-128---127】，编译时不会报错
	//只是转换的结果是按溢出处理，和我们希望的结果不一样
	var num1 int64 =99999
	var num2  = int8(num1)
	fmt.Println("num2=",num2)

}
