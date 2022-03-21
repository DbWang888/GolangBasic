//工厂模式测试

package main
import (
	"fmt"
	"10_factory/model"
)

func main(){
	//var stu model.Student = model.Student{"xiaoming", 12}
	//stu返回的是指针
	stu := model.NewStudent("大明", 15)
	fmt.Println(*stu) //结构体字段首字母小写，可以连同结构体一起输出，
	//但是不能单独访问输出  错误：stu.score
	fmt.Println(stu.GetScore()) 
}