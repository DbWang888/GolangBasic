package main
import (
	"fmt"
)
func test(n int){
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}


//n在1-10天之间
func taozi(n int) int {
	if n == 10 {
		return 1
	} else {
		return (taozi(n + 1) + 1) * 2
	}
}


func main(){
	test(4)
	n := taozi(1)
	fmt.Println(n)
}