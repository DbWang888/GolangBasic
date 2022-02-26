package main
import (
	"fmt"
)

func main(){
	// var ss float64
	// var xb string
	// fmt.Println("输入用时及性别（男或女）")
	// fmt.Scanf("%f %s", &ss, &xb)

	// if ss < 8.0 {
	// 	if xb == "男" {
	// 		fmt.Println("进入男子组")
	// 	} else if xb == "女" {
	// 		fmt.Println("进入女子组")
	// 	}
	// } else {
	// 	fmt.Println("你被淘汰")
	// }
	
	var mon uint8
	var age uint8

	fmt.Println("请输入月份及年龄，空格隔开")
	fmt.Scanf("%d %d", &mon, &age)

	if mon >= 4 && mon <= 10 {
		if age >= 18 && age <= 60 {
			fmt.Println("成人票,票价为60")
		} else if age < 18 {
			fmt.Println("儿童票，票价30")
		} else if age > 60 {
			fmt.Println("老人票，票价20")
		}
	} else {
		if age < 18 || age > 60 {
			fmt.Println("20")
		} else {
			fmt.Println("40")
		}
	}


}