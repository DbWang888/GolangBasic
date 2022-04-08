//并发/并行方式判断0-80000中的哪些数为素数

package main

import (
	"fmt"
	"time"
)

//向intchan管道里面存入1-8000的int类型数
func intNum(intchan chan int) {
	for i := 1; i <= 8000; i++ {
		intchan <- i
	}
	//关闭管道
	close(intchan)
}

//从intchan管道里面取出数并判断是否为素数,
//若为素数则存入isprimechan管道
func isPrime(intchan chan int, isprimechan chan int, exitchan chan bool) {
	// time.Sleep(time.Second)

	var flag bool
	for {
		flag = true //辅助判断是否素数进行存入的量
		num, ok := <-intchan
		if num == 1 {
			flag = false
		}

		if !ok { //如果数被取完 则退出
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明不是素数
				flag = false
				break
			}
		}
		if flag {
			isprimechan <- num
		}
	}

	fmt.Println("有一个协程退出") //此时不能关闭isprimechan管道，只需要向exitchan写入true
	exitchan <- true
}

func main() {
	//建立管道，存入需要判断的数据
	var intChan chan int
	intChan = make(chan int, 1000)

	//建立管道，存入素数
	var isPrimeChan chan int
	isPrimeChan = make(chan int, 3000)

	//建立管道，对是否退出进行判断
	var exitChan chan bool
	exitChan = make(chan bool, 4)

	start := time.Now().UnixNano()
	//开启1个协程，向intchan中放入数据
	go intNum(intChan)

	//开启4个协程，从intchan管道里面取出数并判断是否为素数,
	//若为素数则存入isprimechan管道
	for i := 1; i <= 4; i++ {
		go isPrime(intChan, isPrimeChan, exitChan)
	}

	//判断exitchan是否有4个值，有的话则关闭isprimechan管道
	go func() {
		// for i := 0; i < 4; i++ {
		// 	<-exitChan
		// }
		for {
			if len(exitChan) == 4 {
				close(isPrimeChan)
				end := time.Now().UnixNano()
				fmt.Println("用时 = ", end-start)
				break
			}
		}

	}()

	time.Sleep(time.Second * 2)

	//遍历isprimechan把结果取出
	// for {
	// 	_, ok := <-isPrimeChan
	// 	if !ok {
	// 		break
	// 	}
	// 	//将结果输出
	// 	// fmt.Println("素数 = ", res)
	// }

	fmt.Println("main线程退出")

}
