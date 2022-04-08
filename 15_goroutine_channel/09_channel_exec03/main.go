//开启8个协程对1-2000的每个数进行单独累加，并输出结果
package main

import "fmt"

//将数放入intchan中
func putChan(intChan chan int) {
	for i := 1; i <= 2000; i++ {
		intChan <- i
	}
	close(intChan)
}

//（从intChan中取出数据，进行累加操作，完成后将数放入resChan），
//执行结束则将标识值写入exitChan中
func GetSum(intchan chan int, reschan chan map[int]int, exitchan chan bool) {
	resMap := make(map[int]int, 500)
	for { //
		res := 0
		num, ok := <-intchan
		if !ok {
			break
		}
		for i := 0; i <= num; i++ {
			res += i
		}
		resMap[num] = res
	}
	reschan <- resMap
	exitchan <- true
}

func main() {
	//声明3个chan
	intChan := make(chan int, 200) //存入原始数值
	// resChan := make(chan int, 2000) //存入累加后的值
	resChan := make(chan map[int]int, 8)
	exitChan := make(chan bool, 8) //存入是否结束的标识值

	//开启一个协程将数值放入intChan
	go putChan(intChan)

	//开启8个协程进行累加计算
	for i := 1; i <= 8; i++ {
		go GetSum(intChan, resChan, exitChan)
	}
	//判断是否结束,结束后关闭recChan
	go func() {
		for {
			if len(exitChan) == 8 {
				close(resChan)
				break
			}
		}
	}()

	//遍历resChan输出结果
	for {
		num, ok := <-resChan
		for i, v := range num {
			fmt.Printf("res[%v] = %v\n", i, v)
		}
		if !ok {
			break
		}
	}
	fmt.Println("MAIN线程结束")
}
