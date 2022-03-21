package main

import "fmt"

type A interface {
	test1()
}
type B interface {
	test2()
	test1()
}

type A_B interface {
	A
	B
}
type TestInterface struct {
}

func (t TestInterface) test1() {
	fmt.Println("test1")
}
func (t TestInterface) test2() {
	fmt.Println("test1")
}

type I interface {
}

func main() {
	var n int = 10
	var i I
	i = n

	var n2 float64 = 55.6
	i = n2
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	var t TestInterface
	var ab A_B
	ab = t
	ab.test1()
	fmt.Println(ab)

}
