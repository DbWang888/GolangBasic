package main
import (
	"fmt"
)

//定义计算器结构体
type Calcuator struct {
	Num1 float64
	Num2 float64
}

//定义加减乘除四个方法
func (c *Calcuator) Add() float64 {
	return c.Num1 + c.Num2
}
func (c *Calcuator) Sub() float64 {
	return c.Num1 - c.Num2
}
func (c *Calcuator) Mul() float64 {
	return c.Num1 * c.Num2
}
func (c *Calcuator) Div() float64 {
	return c.Num1 / c.Num2
}
//用一个方法搞定

func (c *Calcuator) CalcuatorFunction (key string) float64 {
	var res float64 = 0
	switch key{
	case "+":
		res = c.Num1 + c.Num2
	case "-":
		res = c.Num1 - c.Num2
	case "*":
		res = c.Num1 * c.Num2
	case "/":
		res = c.Num1 / c.Num2
	default:
		fmt.Println("输入有误")
	}
	return res
}

func main(){
	cal := Calcuator {20.5, 5.8}
	fmt.Println(cal.Div())
	res := cal.CalcuatorFunction("a")
	fmt.Println(res)
}