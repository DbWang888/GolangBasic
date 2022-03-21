package main
import (
	"fmt"
	"strconv"
)

type Visitor struct {
	Name string
	Sex string
	Age int
	Id int
}

func (v *Visitor) getPrice() string {
	str := ""
	if v.Age <= 5 || v.Age >= 60 {
		str = fmt.Sprintf("%s免费", v.Name)
		return str
	} else if v.Age > 5 && v.Age < 18 {
		if v.Sex == "男" {
			str = fmt.Sprintf("%s少年票15元", v.Sex)
			return str
		} else {
			str = fmt.Sprintf("%s少年票10元", v.Sex)
			return str
		}
	} else {
		if v.Sex == "男" {
			str = fmt.Sprintf("%s成人票30元", v.Sex)
			return str
		} else {
			str = fmt.Sprintf("%s成人票20元", v.Sex)
			return str
		}
	}
}

func main(){
	var name, sex,name_id string
	var age,id int
	var viss map[string]Visitor = make(map[string]Visitor,10)
	for {
		fmt.Println("请输入姓名,输入退出则退出")
		fmt.Scanln(&name)
		if name == "退出" {
			fmt.Println("程序退出")
			break
		}
		fmt.Println("请输入性别")
		fmt.Scanln(&sex)
		fmt.Println("请输入年龄")
		fmt.Scanln(&age)
		fmt.Println("请输入ID")
		fmt.Scanln(&id)

		var vis Visitor = Visitor{
			Name:name,
			Sex:sex,
			Age:age,
			Id:id,
		}
		name_id = name + strconv.Itoa(id)
		viss[ name_id ] = vis
		str := vis.getPrice()
		fmt.Println(str)
		fmt.Println(viss)
	}
	

}