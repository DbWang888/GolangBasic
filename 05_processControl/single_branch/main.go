package main
import(
	"fmt"
)

func main(){
	var a int = 0
	fmt.Println("请输入你的年龄：")
	fmt.Scanf("%d", &a)
	if a > 18 {
		fmt.Printf("你已经 %d 岁了", a)
	}
}