package main
import "fmt"

func main(){

	var days int = 97
	var xingqi int 
	var lingjit int
	xingqi = days / 7
	lingjit = days % 7
	fmt.Printf("%d个星期，%d天\n", xingqi, lingjit)

	var huashi float32
	var sheshi float32
	huashi = 134.2

	sheshi = 5.0/9 * (huashi - 100)

	fmt.Println(sheshi)

}