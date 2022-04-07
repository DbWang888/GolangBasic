package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var faceChan chan interface{}
	faceChan = make(chan interface{}, 5)
	xiaoming := Person{"小明", 20}
	faceChan <- xiaoming
	faceChan <- 10.01
	faceChan <- "你好"
	faceChan <- 10

	xiaoming1 := <-faceChan
	//类型断言之后才能使用，如果不用类型断言则为interface{}类型，无法使用
	xiaoming2 := xiaoming1.(Person)

	fmt.Printf("%v的年龄是%v\n", xiaoming2.Name, xiaoming2.Age)

}
