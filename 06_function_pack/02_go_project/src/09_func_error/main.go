package main
import (
	"fmt"
	"errors"
)

func test(){

	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
	}()
	num1 := 100
	num2 := 0
	fmt.Println(num1 / num2)
} 

func test2()(error){
	i := 1
	i2 := 0
	if i == i2 {
		return nil
	} else {
		return errors.New("错误")
	}
}

func main(){
	test()
	fmt.Println("下面代码")

	in := test2()
	panic(in)

	fmt.Println("下面代码111")

}