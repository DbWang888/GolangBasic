package main

import (
	"fmt"
	"time"
	"strconv"
)


func time_test(){
	str := ""
	for i := 0; i <= 100000000; i++ {
		str = "Golang" + strconv.Itoa(i)
	}
	fmt.Println(str)
}


func main(){

	//获取当前时间
	
	now := time.Now()
	fmt.Printf("当前时间是%v\n类型是%T\n", now, now)
	

	//当前月份

	mon := int(now.Month())
	fmt.Printf("当前月份是%v\n类型是%T\n", mon, mon)


	//格式化日期时间  printf（直接打印）  Sprint（返回字符串）

	fmt.Printf("现在时间为：%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	now_time := fmt.Sprintf("现在时间为：%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now_time)

	//格式化日期时间  time.Format()

	fmt.Printf(now.Format("2006-01-02 15:04:05"))

	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// 常量运用，休眠时间

	// for i := 0 ; ; {

	// 	fmt.Printf("%d\n",i)
	// 	time.Sleep(time.Second / 10)
	// 	i++
	// 	if i >= 10 {
	// 		break
	// 	}

	// }
	
	//统计所用时间

	start_time := time.Now().Unix()
	time_test()
	end_time := time.Now().Unix()
	fmt.Printf("所用时间%v", end_time - start_time)



	


}