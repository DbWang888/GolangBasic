package utils
import "fmt"

func GetName() (string, int) {
	var name string 
	var age int
	fmt.Println("请输入名字")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	return name,age
}