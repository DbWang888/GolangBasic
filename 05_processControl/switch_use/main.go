package main
import (
	"fmt"
)

func main(){
	// var char string
	// fmt.Println("输入26字母小写类型")
	// fmt.Scanf("%s", &char)
	// switch {
	// 	case char == "a" :
	// 		fmt.Println("A")
	// 	case char =="b" :
	// 		fmt.Println("B")
		
	// 	case char =="c" :
	// 		fmt.Println("C")
	// 	case char =="d" :
	// 		fmt.Println("D")
		
	// 	case char =="e" :
	// 		fmt.Println("E")
		
	// 	default :
	// 		fmt.Println("other")
		
	// }

	var mon uint8
	fmt.Println("输入月份")
	fmt.Scanf("%d", &mon)

	switch mon {
	case 3, 4, 5 :
		fmt.Println("春季")
	case 6, 7, 8 :
		fmt.Println("夏季")
	case 9, 10, 11 :
		fmt.Println("秋季")
	case 12, 1, 2 :
		fmt.Println("冬季")
	default:
		fmt.Println("ddddd")
	}
	
}