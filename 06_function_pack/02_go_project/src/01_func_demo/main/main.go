package main
import (
	"fmt"
	"01_func_demo/utils"
)

func main(){
	var name string
	var age int
	name, age = utils.GetName()
	fmt.Println(name,age)
}