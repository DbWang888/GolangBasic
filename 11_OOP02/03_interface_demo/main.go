package main

import "fmt"

//定义接口，方法
type Usb interface {
	calssify()
}

//定义电脑结构体
type Computer struct {
}

//定义手机结构体
type Phone struct {
}

//定义手机结构体方法，输出我是手机
func (p Phone) calssify() {
	fmt.Println("我是手机s")
}

//定义相机结构体
type Camera struct {
}

//定义相机方法，输出我是相机
func (c Camera) calssify() {
	fmt.Println("我是相机s")
}

//编写working方法接受USB接口类型变量， usb变量会根据传入的实参来判断到底是phone还是camera
//实现类似于把手机或照相机插入usb接口提示不同的信息的内容
func (c Computer) Working(usb Usb) {
	usb.calssify()
}

func main() {
	var computer Computer = Computer{}
	var camera Camera = Camera{}
	var phone Phone = Phone{}

	computer.Working(camera)
	computer.Working(phone)

}
