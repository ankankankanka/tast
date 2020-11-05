# tast
项目描述
package main
import (
	"fmt"
)
//演示golang中基本数据练习转成string使用
func main() {

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar bute = 'h'
	var str string //空的str

	//使用第一种方式转换fmt.sprintf 方法
	str = fmt.Printf("%d",num1)
	fmt.Printf("str type %t str=%v",str ,str)

	str = fmt.Sprint("%f",num2)
	fmt.Printf("str type %T str=%v\n",str ,str)

	str = fmt.Sprintf("%b",num2)


}
