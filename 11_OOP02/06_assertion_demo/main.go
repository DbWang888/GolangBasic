package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	//类型断言案例1
	var a interface{}
	var p Point = Point{1, 2}
	a = p
	var b Point
	//b = a会报错，这里需要类型断言
	//b = a.(Point) 意思是 a是否可以转换成Point类型，如果是则赋值给b
	b = a.(Point)
	fmt.Println(b)

	//类型断言案例2
	var x interface{} //空接口，空接口可以接收任意类型
	var b2 float32 = 1.2
	x = b2
	y := x.(float32)

	fmt.Printf("类型是%T\n", y)

	//带检测的类型断言--安全写法  不会报panic
	var x2 interface{}
	var b3 float32 = 2.3
	x2 = b3

	y2, ok := x2.(float64)
	if ok == true {
		fmt.Println(y2)
	} else {
		fmt.Println("断言失败")
	}

}
