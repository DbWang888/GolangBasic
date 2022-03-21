package main
import (
	"fmt"
)

type Person struct {
	Name string
	Skill string
}

type Circle struct {
	radius float64
}

//Person类型的方法，只能通过Person类型的变量来调用
func (p_name Person) print_myname() {
	fmt.Println(p_name.Name)
}

//输出xxx是一个好人
func (person Person) speak(){
	fmt.Printf("%v是一个好人\n", person.Name)
}

//输出1-1000的和
func (person Person) jisuan(){
	sum := 0
	for i:= 1; i <= 1000; i++ {
		sum = sum + i
	}
	fmt.Printf("1-1000的和为%d\n ", sum)
}

//接收一个数，输出1到n的和
func (person Person) jisuan2 (n int){
	sum := 0
	for i:= 1; i <= n; i++ {
		sum = sum + i
	}
	fmt.Printf("1-n的和为%d\n", sum)
}

//计算两个数的和，并返回结果
func (person Person) getSum (n1 int, n2 int)( int ){
	sum := n1 + n2
	return sum
}

//返回圆的面积
func (cir Circle) area()( float64 ){
	return 3.14 * cir.radius * cir.radius
}

func main(){
	jack := Person {"jack", "打印"}
	jack.print_myname()
	jack.speak()
	jack.jisuan()
	jack.jisuan2(100)
	sum := jack.getSum(10, 20)
	fmt.Printf("和为%d\n", sum)

	cir1 := Circle{3}
	fmt.Println(cir1.area())

}