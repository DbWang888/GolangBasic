package main

import "fmt"

//-------------------类型断言的最佳实践1-----------------------
//定义Usb接口，实现类似于电脑插入不同设备的不同反应
//Phone结构体增加特有方法Call(),当Usb接口接收的是Phone类型变量时，调用Call()方法
//声明Usb接口
type Usb interface {
	start()
	end()
}

type Phone struct {
}

func (p Phone) start() {
	fmt.Println("手机开始工作")
}

func (p Phone) end() {
	fmt.Println("手机结束工作")
}

func (p Phone) Call() {
	fmt.Println("手机打电话")
}

type Camera struct {
}

func (c Camera) start() {
	fmt.Println("照相机开始工作")
}

func (c Camera) end() {
	fmt.Println("照相机结束工作")
}

type Computer struct {
}

//定义Computer方法，根据传入的usb接口实例不同，调用不同结构体的方法
func (c Computer) working(usb Usb) {
	usb.start()
	if usb, ok := usb.(Phone); ok {
		usb.Call()
	}
	usb.end()
}

//-------------------类型断言的最佳实践1-----------------------

//-------------------类型断言的最佳实践2-----------------------

type Student struct {
	Name string
}

//写一个函数，循环判断传入的参数类型
func TypeJudge(items ...interface{}) {
	fmt.Printf("%T\n", items) //items为空接口切片 []interface{}
	for i, x := range items {
		switch x.(type) { //这里type为关键字，固定写法
		case int, int32, int64:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		case string:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		case float32:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		case bool:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		case Student:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		case *Student:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		case nil:
			fmt.Printf("序列号是%d, 值是%v\n", i, x)
		default:
			fmt.Printf("序列号%d的类型未知，值为%v\n", i, x)
		}

	}
}

//-------------------类型断言的最佳实践2-----------------------

func main() {
	//-------------------类型断言的最佳实践1-----------------------
	var computer Computer
	var phone Phone
	var camera Camera
	computer.working(phone)
	computer.working(camera)
	//-------------------类型断言的最佳实践1-----------------------
	fmt.Println()
	fmt.Println()
	fmt.Println()

	//-------------------类型断言的最佳实践2-----------------------
	var a int = 66
	var b float32 = 33.33
	var c string = "wodetian"
	var d bool = true
	var e int8 = 1
	var f Student = Student{"小明"}
	var g *Student = &f
	TypeJudge(a, b, c, d, e, f, g)
	//-------------------类型断言的最佳实践2-----------------------
}
