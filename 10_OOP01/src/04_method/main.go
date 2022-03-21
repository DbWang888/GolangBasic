package main
import (
	"fmt"
)
//int和float32也有对应的方法
type integer int

func (i *integer) print (){
	fmt.Println("i = ", *i)
	*i = *i + 20
}


//给 *student 实现string方法
type Student struct {
	Name string
	Age int
}
func (s *Student) String() string {
	str := fmt.Sprintf("Name = %v , Age = %v", s.Name, s.Age)
	return str
}

func main(){
	var i integer = 10
	i.print()
	fmt.Println("i2 = ", i)
	stu := Student{"jack", 20}
	//如果实现了string方法，就会自动调用，且方法名大写
	fmt.Println(&stu)
}