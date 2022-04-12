package main

import (
	"fmt"
	"reflect"
)

//对基本数据类型反射
func testReflect1(b interface{}) {
	//获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("typeOf = %v\n", rType)

	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("ValueOf = %v\n", rVal)

	//返回整数
	intval := rVal.Int() + 2
	fmt.Printf("intval = %v  Type = %T\n", intval, intval)

	//将reflect.Value转成interfacle{}，再通过类型断言转成需要类型
	val := rVal.Interface()
	i2 := val.(int)
	fmt.Printf("i2 = %v  i2 = %T\n", i2, i2)
}

//对结构体反射
func testReflect2(b interface{}) {
	//获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("typeOf = %v\n", rType)

	//获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("ValueOf = %v\n", rVal)

	//获取kind
	kind1 := rType.Kind()
	kind2 := rVal.Kind()
	fmt.Printf("kind1 = %v  kind2 = %v \n", kind1, kind2)

	//将reflect.Value转成interfacle{}，再通过类型断言转成需要类型
	val := rVal.Interface()
	//此处类型为编译中有效，所以即使val也为Person类型，但是无法使用Val.Name
	fmt.Printf("val = %v  val = %T\n", val, val)
	stu, ok := val.(Person)
	if ok {
		fmt.Printf("stuNAME = %v  stu = %T\n", stu.Name, stu)
	}

}

type Person struct {
	Name string
	Age  int
}

func main() {
	var i int = 100
	testReflect1(i)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	person1 := Person{
		Name: "小明",
		Age:  20,
	}
	testReflect2(person1)
}
