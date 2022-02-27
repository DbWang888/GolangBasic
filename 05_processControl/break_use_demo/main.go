package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var count int = 0
	rand.Seed(time.Now().Unix())
	for {
	
		num := rand.Intn(101)
		//  fmt.Println("num = ", num)
		count++
		if num == 99 {
			break
		}
	}
	fmt.Println("用的次数为",count)

	//标签使用
	label:
	for i := 0; i <= 5; i++ {
		for j := 0;j <= 3; j++ {
			fmt.Println(j)
			if j == 2 {
				break label
			}
		}
	}

	//登录验证，提示还剩几次机会

	var jihui int = 3
	var usename string
	var pwd string

	for i := 1; i <= jihui; i++ {
		fmt.Println("请输入账号")
		fmt.Scanln(&usename)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if usename == "张无忌" && pwd == "abc123" {
			fmt.Println("ok")
			break
		} else {
			if jihui - i == 0 {
				fmt.Printf("失败")
			} else {
				fmt.Printf("还有%d次机会", jihui - i)
			}
			
		}
		
	}
	



}