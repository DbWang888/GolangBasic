package main
import (
	"fmt"
)

//指针传递
func change(p *int){
	*p = 50
}

//函数也是一种数据类型
func getSum(n1 int, n2 int)(int){
	return n1 + n2
}

//自定义函数类型
type myFuncType func (int, int) int

func getSum2(myfunc myFuncType, n1 int, n2 int) int {
	return myfunc(n1, n2)
}

//对返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//可变参数
func sum1 (n1 int,args... int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}


func main(){
	var i = 10
	change(&i)
	fmt.Println(i)
	
	a := getSum
	res := a(1,1)
	fmt.Printf("a的类型为%T\ngetSum的类型为%T\nres=%d\n", a, getSum, res)
	//自定义函数类型

	type myInt int
	var num myInt
	fmt.Printf("num类型%T\n",num)

	num2 := getSum2(getSum, 10, 20)
	fmt.Println("num2 = ", num2)

	c, d := getSumAndSub(10,20)
	fmt.Println(c,d)

	sumsum := sum1(1,1,1,1,1)
	fmt.Println(sumsum)
}