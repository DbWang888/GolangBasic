package main
import (
	"fmt"
	"strings"
)
//累加器
func adder()(func (int) int) {
	var n int  //  n在返回的匿名函数外时，相当于匿名函数调用外域n,可以实现累加
	return func (x int) int {
		//var n int  //n在返回的匿名函数内时 ，则不会累加
		n += x
		return n
	}
}

//判断后缀
func makeSuffix(name string)(func (string) string) {

	return func (names string) string {
		if strings.HasSuffix(names, name) == true {
			return names
		} 
		return names + name
		
	}
}


//文件后缀

func main(){
	//累加器
	add := adder()
	xxx := add(5)
	fmt.Println(xxx)

	//判断后缀
	houzhui := ".png"
	if_name := makeSuffix(houzhui)
	dir_name := "wode.jpg"
	fanhui := if_name(dir_name)
	fmt.Println(fanhui)

	
	
}