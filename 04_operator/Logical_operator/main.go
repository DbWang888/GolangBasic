package main
import "fmt"

func test() bool {

	fmt.Println("123")
	return true

}

func main(){

	a := true
	b := false

	fmt.Println(a && test())   //短路与
	fmt.Println(b || test())  //短路或
	fmt.Println(b && test())	//短路与  前面为false  不执行test()
	fmt.Println( ! (a && b) )
}