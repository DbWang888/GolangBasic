package main
import "fmt"
import "math"

func main(){
	//案例1
	// var num float64
	// fmt.Println("请输入成绩：")
	// fmt.Scanln(&num)

	// if num == 100.0 {
	// 	fmt.Println("奖励BWN")
	// } else if num > 80.0 && num <= 99 {
	// 	fmt.Println("奖励ip7")
	// } else if num > 60 && num <= 80 {
	// 	fmt.Println("奖励ipad")
	// } else {
	// 	fmt.Println("什么也没有")
	// }

	//案例2
	var a float64
	var b float64
	var c float64
	var sqlabc float64

	fmt.Println("请输入三个数字，a b c，以空格隔开")
	fmt.Scanf("%f %f %f", &a, &b, &c)

	sqlabc = b * b - 4 * a * c

	if sqlabc > 0 {
		var x1 , x2 float64
		x1 = (-b + math.Sqrt(sqlabc)) / (2 * a)
		x2 = (-b - math.Sqrt(sqlabc)) / (2 * a)
		fmt.Printf("函数有两个解，x1 = %f ，x2 = %f", x1, x2)
	} else if sqlabc == 0 {
		var x1 float64
		x1 = (-b + math.Sqrt(sqlabc)) / (2 * a)
		fmt.Printf("函数只有一个解，x1=x2=%f", x1)
	} else if sqlabc < 0 {
		fmt.Println("函数无解")
	}

}