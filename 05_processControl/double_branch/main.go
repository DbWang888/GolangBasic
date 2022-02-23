package main
import(
	"fmt"
)

func main(){
	// var a int = 0
	// fmt.Println("请输入你的年龄：")
	// fmt.Scanf("%d", &a)
	// if a > 18 {
	// 	fmt.Printf("你已经 %d 岁了", a)
	// } else {
	// 	fmt.Printf("你还小")
	// }

	// var n1 int32 = 40
	// var n2 int32 = 24
	// if n1 + n2 >= 50 {
	// 	fmt.Printf("hello word\n")
	// }
	
	// var n3 float64 = 33.3
	// var n4 float64 = 15.5

	// if n3 > 10 && n4 < 20 {
	// 	sum := n3 + n4
	// 	fmt.Printf("两个数的和为：%f\n", sum)
	// }

	var num1 int32 = 15
	var num2 int32 = 3
	if (num1 + num2) % 3 == 0 && (num1 + num2) % 5 == 0 {
		fmt.Printf("二者的和可以被3和5整除")
	} 
	if (num1 + num2) % 3 != 0 && (num1 + num2) % 5 == 0 {
		fmt.Printf("二者的和可以被5整除")
	} 
	if (num1 + num2) % 3 == 0 && (num1 + num2) % 5 != 0 {
		fmt.Printf("二者的和可以被3整除")
	} 

	


}