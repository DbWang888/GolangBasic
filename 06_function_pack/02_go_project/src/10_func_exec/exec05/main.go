package main
import (
	"fmt"
	"time"
)


//得到当前时间

func get_now_time()(string){
	now := time.Now()
	now_time := now.Format("2006-01-02 15:04:15")
	return now_time
}

func win(now){
	fmt.Println("--------小小计算器--------")
	fmt.Println("当前时间：",now)
	fmt.Println("1、加法")
	fmt.Println("2、减法")
	fmt.Println("3、乘法")
	fmt.Println("4、除法")
	fmt.Println("0、结束")
}

func main(){

}