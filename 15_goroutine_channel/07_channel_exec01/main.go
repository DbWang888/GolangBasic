//并发/并行方式判断0-80000中的哪些数为素数

package main

import (
	"fmt"
)

//判断一个数是否是素数
func ifPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//将数写入管道
func writeChan(inchan chan int, n1 int, n2 int) {
	for i := n1; i <= n2; i++ {
		inchan <- i
	}
	close(inchan)
}

//从管道中取出数，并判断是否是素数,将素数输出,放入切片
func GetPrimeChan(inchan chan int, bochan chan bool, isP *[]int) {
	b := true

	for i := range inchan {
		b = ifPrime(i)
		if b == true {
			*isP = append(*isP, i)
			fmt.Printf("素数 = %d\n", i)
		}
	}
	bochan <- true
	close(bochan)
}

// //将存管道中的素数输出
// func outChanPr(prchan chan int, bochan chan bool){
// 	for i := range prchan
// }

func main() {

	//定义管道
	var inchan1 chan int
	inchan1 = make(chan int, 200)
	var inchan2 chan int
	inchan2 = make(chan int, 200)
	var bochan chan bool
	bochan = make(chan bool, 2)
	var prim1 []int
	prim1 = make([]int, 20)
	var prim2 []int
	prim2 = make([]int, 20)
	var prim3 []int
	prim3 = make([]int, 40)

	go writeChan(inchan1, 1, 40000)
	go writeChan(inchan2, 40001, 80000)
	go GetPrimeChan(inchan1, bochan, &prim1)
	go GetPrimeChan(inchan2, bochan, &prim2)

	for {
		_, ok := <-bochan
		if ok == false {
			break
		}
	}
	prim3 = append(prim3, prim1...)
	prim3 = append(prim3, prim2...)
	fmt.Println(prim3)
	fmt.Println()
	fmt.Println()
	fmt.Println(len(prim3))

}
