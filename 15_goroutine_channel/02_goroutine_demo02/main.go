package main

import (
	"fmt"
	"sync"
	"time"
)

//全局变量互斥锁  解决同步问题

//声明全局变量
var (
	//存放数据
	myMap = make(map[int]int, 10)

	//锁 sync包的Mutex
	lock sync.Mutex
)

//sum计算累加到n的和，将这个结果放到Mymap中

func Sum(n int) {
	res := 1
	for i := 0; i <= n; i++ {
		res += i
	}
	//这里将res放入到myMap中
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {

	//开启20个协程完成
	for i := 1; i <= 50; i++ {
		go Sum(i)
	}

	//休眠 10分钟 再遍历map
	time.Sleep(time.Second * 10)

	//此处必须加锁，否则还会存在资源竞争
	//因为主线程和协程未通信，主线程无法检测协程是否运行结束
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
