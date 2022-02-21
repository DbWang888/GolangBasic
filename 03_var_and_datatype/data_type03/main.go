package main
import (
	"fmt"
)

//指针
func main(){
	var num int = 10
	fmt.Println("num的地址为",&num)

	var pnum *int = &num
	fmt.Println("num的值",*pnum)
	fmt.Println("pnum的地址",&pnum)

	*pnum = 15
	fmt.Println("num的值",num)


	

}