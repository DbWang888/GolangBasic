package main

import (
	"fmt"
)

func main() {
	var intChan chan int         //声明一个管道变量
	intChan = make(chan int, 10) //初始化管道变量
	intChan <- 50                //向管道中添加数据
	fmt.Printf("len = %v, cap = %v\n", len(intChan), cap(intChan))
	n1 := <-intChan //向管道中取出数据
	fmt.Println(n1)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	n2 := 0
	for i := 0; i < 10; i++ {
		n2 = <-intChan
		fmt.Println(n2)
	}

}
