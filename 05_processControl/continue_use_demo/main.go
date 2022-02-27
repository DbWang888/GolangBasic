package main
import "fmt"

func main(){
	// label:
	// for i := 0; i <= 5; i++ {

	// 	for j := 0; j <= 5; j++ {
			
	// 		if j == 2 {
	// 			continue label
	// 		}
	// 		fmt.Println(i,j)
	// 	}
	// }

	//1-100的奇数
	// for i := 1; i <= 100; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//输入正负数
	// var num int64
	// var count_zs, count_fs int64 
	// for {
	// 	fmt.Println("请输入数据")
	// 	fmt.Scanln(&num)
	// 	if num == 0 {
	// 		fmt.Println(count_zs,count_fs)
	// 		break
	// 	} else {
	// 		if num > 0 {
	// 			count_zs++
	// 			// continue
	// 		} else if num < 0 {
	// 			count_fs++
	// 			// continue
	// 		}
	// 	}
		
	// }

	//判断次数
	var money float64 = 100000.0
	var count int64 = 0
	for {
		if money > 50000.0 {
			money = money * 0.95
			count++
			fmt.Println("1")
			continue
		}
		if money <= 50000.0 && money >= 1000.0 {
			money = money - 1000.0
			count++
			fmt.Println(money)
			continue
		}
		
		if money < 1000.0 {
			fmt.Println("次数为",count)
			break
		}
	}
}