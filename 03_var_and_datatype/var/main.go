package main
import "fmt"

//全局变量声明方式

var (
	a = 1
	b = 2
	c = 3
)


func main(){
	//单个声明变量的三种方法
	fmt.Println("单个变量声明")
	//第一种：同时声明变量名称、变量类型，使用方法：var [变量名称] [变量类型]
	var i int  //整数类型默认值为0
	i = 1000000
	fmt.Println("第一种方式：i = ", i)

	//第二种：根据值自行推断变量类型（类型推导）使用方法：var m = 10.22，m为float类型
	var m = 10.11
	fmt.Println("第二种方式: m = ", m)

	//第三种：省略var 使用方法： [变量名称] := [值]

	n := 10
	fmt.Println("第三种方式：n = ", n)

	//多个变量声明的三种方式
	fmt.Println("多个变量声明")
	//第一种：同时声明变量名称，变量类型，使用方法： var n1, n2, n3 int
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//第二种：类型推导  var n1, n2, n3 = 1,"id",3
	var n4, n5, n6 = 100, "name", 200
	fmt.Println(n4, n5, n6)
	//第三种：类型推导 n1, n2, n3 := "1", 2, 3
	n7, n8, n9 := "1", 2, 3
	fmt.Println(n7, n8, n9)

	//输出全局变量
	fmt.Println("全局变量")
	fmt.Println("全局变量abc", a, b, c)



	

}