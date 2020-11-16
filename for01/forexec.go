package main
import "fmt"

func main () {
	//统计三个班的及格人数，每个班有5名同学
	//分析思路
	//走代码实现
	var classnum int = 2
	var stunum int = 5
	var totalsum float64 = 0.0
	var passcount int = 0
	for j := 1; j <= classnum; j ++ {
	sum := 0.0
	for i := 1; i <= stunum; i++ {
		var score float64
		fmt.Printf("请输入第%d班 第%d个学生的成绩 \n", j, i)
		fmt.Scanln(&score)
		//累计总分
		sum += score
		//判断分数是否及格
		if score >= 60 {
			passcount++
		}
	}
	fmt.Printf("第%d个班级的平均分是%v \n",totalsum, totalsum / float64(stunum * classnum))
	fmt.Printf("及格人数为%\n",passcount)
	}