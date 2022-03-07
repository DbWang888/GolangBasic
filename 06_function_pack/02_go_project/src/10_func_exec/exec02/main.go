package main
import (
	"fmt"
	"time"
	"math/rand"
)

func get_int()(int){
	var n int = 0
	rand.Seed(time.Now().Unix())
	n = rand.Intn(101)
	fmt.Println(n)
	return n
}


func main(){
	var i int = 0
	var num1 int
	num2 := get_int()
	for {
		i++
		fmt.Println("请输入猜的数")
		fmt.Scanln(&num1)
		if num1 == num2 {
			if i == 1 {
				fmt.Println("猜对了，你真的是个天才")
			} else if i > 1 && i < 4 {
				fmt.Println("猜对了，快赶上我了")
			} else if i >= 4 && i <= 9 {
				fmt.Println("一般般")
			} else if i == 10 {
				fmt.Println("最后一次")
			} else {
				fmt.Println("说你点啥呢")
			}
			break
		} else {
			continue
		}
		
	}
	


}