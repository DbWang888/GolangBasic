package main
import (
	"fmt"
)

func sum (n1 int, n2 int) int {
	//defer后的语句连同对应的变量值拷贝入栈，所以defer的值不受n1++影响
	defer fmt.Println("n1", n1)
	defer fmt.Println("n2", n2)
	n1++
	n2++
	fmt.Println(n1,n2)
	return n1 + n2
}


func main(){
	res := sum(5, 6)
	fmt.Println("res",res)
}