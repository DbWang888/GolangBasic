package main

import (
	"fmt"
	"reflect"
)

//反射修改原变量值
func testRef(b interface{}) {

	rVal := reflect.ValueOf(b)
	fmt.Printf("rval = %v, rvalType = %T\n", rVal, rVal)

	rVal.Elem().SetInt(200)

}

func main() {
	var i int = 100
	//涉及到修改变量值，此处应传入指针，再使用rval.Elem()转换成值，使用SetInt()修改
	testRef(&i)
	fmt.Println("i = ", i)
}
