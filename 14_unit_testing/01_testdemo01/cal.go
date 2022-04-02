package testdemo01

import "fmt"

//关于计算的函数

func Add(n1, n2 int) int {
	fmt.Println("run Add")
	return n1 + n2
}

func GetSum(n int) int {
	fmt.Println("run GetSum")
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}
