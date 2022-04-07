package main

import (
	"fmt"
	"time"
)

func send() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello word %d\n", i)
		time.Sleep(time.Second)
	}

}

func main() {

	//开启一个协程
	go send()

	for i := 0; i < 10; i++ {
		fmt.Printf("hello main %d\n", i)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		fmt.Println()
	}

}
