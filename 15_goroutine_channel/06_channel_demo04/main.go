package main

import (
	"fmt"
)

//向管道中写入50个数据
func WriteChan(inchan chan int) {
	for i := 1; i <= 50; i++ {
		inchan <- i
		fmt.Printf("写入第%d个值，为%v\n", i, i)
	}
	//写入完毕后关闭，不影响读取
	close(inchan)
}

//从管道中读取数据，循环遍历输出
func readChan(inchan chan int, bochan chan bool) {
	for i := range inchan {
		fmt.Println("读取的值 = ", i)
	}
	bochan <- true
	close(bochan)
}

func main() {

	//创建两个管道
	var intChan chan int
	intChan = make(chan int, 10)
	var bochan chan bool
	bochan = make(chan bool, 1)

	//开启协程
	go WriteChan(intChan)
	go readChan(intChan, bochan)

	//判断结束主线程
	for {
		i, ok := <-bochan
		if ok == false {
			fmt.Println(i, 111)
			break
		}
		fmt.Println(i, 222)
	}

}
