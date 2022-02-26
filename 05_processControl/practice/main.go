package main
import (
	"fmt"
)

func main(){
	//练习1 输入和判断张三

	// var get_use_name string
	// var get_password string
	// fmt.Println("请输入用户名和密码：")
	// fmt.Scanf("%s %s", &get_use_name, &get_password)

	// if get_use_name == "张三" && get_password == "aa12345" {
	// 	fmt.Println("登录成功")
	// } else {
	// 	fmt.Println("登录失败")
	// }

	//switch 根据月份和年份，求出该月的天数
	var month, year int64
	fmt.Println("请输入年份和月份：")
	fmt.Scanf("%d %d", &year, &month)
	
	// if year % 4 == 0 {
	// 	switch month {
	// 	case 1,3,5,7,8,10,12 :
	// 		fmt.Println("31天")
	// 	case 2 :
	// 		fmt.Println("29天")
	// 	default:
	// 		fmt.Println("30天")
	// 	}

	// } 
	switch month {
	case 1,3,5,7,8,10,12 :
		fmt.Println("31天")
	case 2 :
		if year % 4 == 0 {
			fmt.Println("29天")
		} else {
			fmt.Println("28天")
		}
		
	default:
		fmt.Println("30天")
			}
	
}