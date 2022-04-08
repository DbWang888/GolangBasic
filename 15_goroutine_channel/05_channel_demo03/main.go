package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 200)
	//向管道中放入200个数据
	for i := 1; i <= 200; i++ {
		intChan <- i
	}

	//关闭管道后不影响读取，但是无法写入
	//关闭管道后读取才不会报错--deadlock
	close(intChan)

	//for-range遍历管道,遍历管道前需先close
	for v := range intChan {
		fmt.Println("intChan数据 = ", v)
	}

}
