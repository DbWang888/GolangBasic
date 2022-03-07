package main
import "fmt"

func main(){
	//new用来分配内存，主要用于值类型
	num1 := new(int)
	fmt.Printf("类型%T\n值%v\n地址%v\n", num1, num1, &num1)
}